from __future__ import annotations

import json
import sys
import urllib.error
import urllib.parse
import urllib.request
from dataclasses import dataclass
from datetime import datetime
from typing import Any, Final

DEFAULT_BASE_URL: Final[str] = "https://pilot.example.invalid/vrp/api/v1"
DEFAULT_TIMEOUT_SECONDS: Final[float] = 10.0
MAX_RESPONSE_BYTES: Final[int] = 1024 * 1024


class PublicAPIError(RuntimeError):
    def __init__(
        self,
        code: str,
        message: str,
        request_id: str | None = None,
    ) -> None:
        super().__init__(f"{code}: {message}")
        self.code = code
        self.message = message
        self.request_id = request_id


@dataclass(frozen=True)
class PublicHealth:
    schema: str
    status: str
    observed_at: datetime
    boundary_version: str


@dataclass(frozen=True)
class PublicCapabilities:
    schema: str
    boundary_version: str
    capabilities: tuple[str, ...]


class VRPPublicClient:
    def __init__(
        self,
        base_url: str = DEFAULT_BASE_URL,
        timeout_seconds: float = DEFAULT_TIMEOUT_SECONDS,
    ) -> None:
        normalized = base_url.strip().rstrip("/")

        if not normalized:
            raise ValueError("base_url is required")

        if timeout_seconds <= 0:
            raise ValueError("timeout_seconds must be positive")

        self._base_url = normalized
        self._timeout_seconds = timeout_seconds

    def get_health(self) -> PublicHealth:
        payload = self._get_json("/health")

        schema = require_string(payload, "schema")
        if schema != "vrp.health/v1":
            raise ValueError(f"unexpected health schema: {schema!r}")

        return PublicHealth(
            schema=schema,
            status=require_string(payload, "status"),
            observed_at=parse_datetime(
                require_string(payload, "observed_at")
            ),
            boundary_version=require_string(
                payload,
                "boundary_version",
            ),
        )

    def get_capabilities(self) -> PublicCapabilities:
        payload = self._get_json("/capabilities")

        schema = require_string(payload, "schema")
        if schema != "vrp.capabilities/v1":
            raise ValueError(
                f"unexpected capabilities schema: {schema!r}"
            )

        raw_capabilities = payload.get("capabilities")
        if not isinstance(raw_capabilities, list):
            raise ValueError("capabilities must be an array")

        capabilities: list[str] = []

        for item in raw_capabilities:
            if not isinstance(item, str) or not item:
                raise ValueError(
                    "every capability must be a non-empty string"
                )

            capabilities.append(item)

        return PublicCapabilities(
            schema=schema,
            boundary_version=require_string(
                payload,
                "boundary_version",
            ),
            capabilities=tuple(capabilities),
        )

    def _get_json(self, path: str) -> dict[str, Any]:
        url = self._build_url(path)

        request = urllib.request.Request(
            url=url,
            method="GET",
            headers={
                "Accept": "application/json",
                "User-Agent": "vrp-public-boundary-python-example/0.1",
            },
        )

        try:
            with urllib.request.urlopen(
                request,
                timeout=self._timeout_seconds,
            ) as response:
                body = response.read(MAX_RESPONSE_BYTES + 1)

                if len(body) > MAX_RESPONSE_BYTES:
                    raise ValueError(
                        "public API response exceeded size limit"
                    )

                return decode_json_object(body)

        except urllib.error.HTTPError as exc:
            body = exc.read(MAX_RESPONSE_BYTES + 1)

            try:
                payload = decode_json_object(body)
            except ValueError:
                raise PublicAPIError(
                    code="HTTP_ERROR",
                    message=f"public API returned HTTP {exc.code}",
                ) from exc

            raise PublicAPIError(
                code=str(payload.get("code", "UNKNOWN_ERROR")),
                message=str(
                    payload.get(
                        "message",
                        f"public API returned HTTP {exc.code}",
                    )
                ),
                request_id=optional_string(
                    payload.get("request_id")
                ),
            ) from exc

        except urllib.error.URLError as exc:
            raise ConnectionError(
                f"public API request failed: {exc.reason}"
            ) from exc

    def _build_url(self, path: str) -> str:
        if not path.startswith("/"):
            raise ValueError("path must start with '/'")

        return self._base_url + path


def decode_json_object(data: bytes) -> dict[str, Any]:
    try:
        decoded = json.loads(data.decode("utf-8"))
    except (UnicodeDecodeError, json.JSONDecodeError) as exc:
        raise ValueError("response is not valid UTF-8 JSON") from exc

    if not isinstance(decoded, dict):
        raise ValueError("response JSON must be an object")

    return decoded


def require_string(
    payload: dict[str, Any],
    field: str,
) -> str:
    value = payload.get(field)

    if not isinstance(value, str) or not value:
        raise ValueError(
            f"{field!r} must be a non-empty string"
        )

    return value


def optional_string(value: Any) -> str | None:
    if value is None:
        return None

    if isinstance(value, str):
        return value

    return str(value)


def parse_datetime(value: str) -> datetime:
    normalized = value.replace("Z", "+00:00")

    try:
        return datetime.fromisoformat(normalized)
    except ValueError as exc:
        raise ValueError(
            f"invalid RFC 3339 timestamp: {value!r}"
        ) from exc


def main() -> int:
    base_url = (
        sys.argv[1]
        if len(sys.argv) > 1
        else DEFAULT_BASE_URL
    )

    client = VRPPublicClient(base_url=base_url)

    try:
        health = client.get_health()

        print(f"SCHEMA={health.schema}")
        print(f"STATUS={health.status}")
        print(f"BOUNDARY_VERSION={health.boundary_version}")
        print(
            "OBSERVED_AT="
            + health.observed_at.isoformat()
        )

        capabilities = client.get_capabilities()

        print(f"CAPABILITIES_SCHEMA={capabilities.schema}")
        print(
            "CAPABILITIES_BOUNDARY_VERSION="
            + capabilities.boundary_version
        )

        for capability in capabilities.capabilities:
            print(f"CAPABILITY={capability}")

        return 0

    except (
        ConnectionError,
        PublicAPIError,
        ValueError,
    ) as exc:
        print(
            f"PUBLIC_API_CLIENT_FAILED={exc}",
            file=sys.stderr,
        )
        return 1


if __name__ == "__main__":
    raise SystemExit(main())