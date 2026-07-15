from dataclasses import dataclass
from datetime import datetime
from typing import Optional


@dataclass(slots=True)
class PublicHealth:
    schema: str
    status: str
    observed_at: datetime
    boundary_version: str


@dataclass(slots=True)
class PublicCapabilities:
    schema: str
    boundary_version: str
    capabilities: list[str]


@dataclass(slots=True)
class EvidenceReference:
    schema: str
    evidence_id: str
    digest: str
    created_at: datetime
    media_type: str
    retrieval_status: str


@dataclass(slots=True)
class ValidationStatus:
    schema: str
    request_id: str
    status: str
    updated_at: datetime
    verdict_available: bool


@dataclass(slots=True)
class PublicVerdict:
    schema: str
    request_id: str
    verdict: str
    reason_code: str
    completed_at: datetime
    evidence: Optional[EvidenceReference]