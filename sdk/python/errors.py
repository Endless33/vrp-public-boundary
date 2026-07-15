class VRPPublicError(Exception):
    """Base exception for the VRP Public Engineering SDK."""


class APIError(VRPPublicError):
    def __init__(
        self,
        code: str,
        message: str,
        request_id: str | None = None,
    ) -> None:
        self.code = code
        self.message = message
        self.request_id = request_id

        text = f"{code}: {message}"

        if request_id:
            text += f" (request_id={request_id})"

        super().__init__(text)