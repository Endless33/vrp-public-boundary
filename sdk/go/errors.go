package vrpboundary

import "fmt"

// ErrorCode is a stable public API error classification.
type ErrorCode string

const (
	ErrorInvalidRequest    ErrorCode = "INVALID_REQUEST"
	ErrorUnauthorized      ErrorCode = "UNAUTHORIZED"
	ErrorNotFound          ErrorCode = "NOT_FOUND"
	ErrorConflict          ErrorCode = "CONFLICT"
	ErrorRateLimited       ErrorCode = "RATE_LIMITED"
	ErrorServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"
	ErrorUnexpectedResponse ErrorCode = "UNEXPECTED_RESPONSE"
	ErrorTransportFailure   ErrorCode = "TRANSPORT_FAILURE"
)

// APIError represents an error returned through the public engineering
// boundary.
//
// It intentionally contains no Protected Runtime state, private stack traces,
// internal identifiers, or decision-engine details.
type APIError struct {
	Schema     string    `json:"schema"`
	Code       ErrorCode `json:"code"`
	Message    string    `json:"message"`
	RequestID  string    `json:"request_id,omitempty"`
	HTTPStatus int       `json:"-"`
	Cause      error     `json:"-"`
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e == nil {
		return "<nil>"
	}

	if e.RequestID != "" {
		return fmt.Sprintf(
			"vrp public API error: code=%s status=%d request_id=%s message=%s",
			e.Code,
			e.HTTPStatus,
			e.RequestID,
			e.Message,
		)
	}

	return fmt.Sprintf(
		"vrp public API error: code=%s status=%d message=%s",
		e.Code,
		e.HTTPStatus,
		e.Message,
	)
}

// Unwrap exposes the underlying transport or decoding error when present.
func (e *APIError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Cause
}

// IsCode reports whether err is an APIError with the requested public code.
func IsCode(err error, code ErrorCode) bool {
	apiErr, ok := err.(*APIError)
	if !ok {
		return false
	}

	return apiErr.Code == code
}