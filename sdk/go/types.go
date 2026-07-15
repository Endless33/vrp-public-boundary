package vrpboundary

import "time"

// HealthStatus describes bounded public availability.
//
// It does not expose internal runtime state, topology, protected counters,
// customer data, or private telemetry.
type HealthStatus string

const (
	HealthStatusAvailable   HealthStatus = "AVAILABLE"
	HealthStatusDegraded    HealthStatus = "DEGRADED"
	HealthStatusUnavailable HealthStatus = "UNAVAILABLE"
)

// ValidationState describes the externally observable lifecycle of a
// validation request.
type ValidationState string

const (
	ValidationStateAccepted  ValidationState = "ACCEPTED"
	ValidationStateRunning   ValidationState = "RUNNING"
	ValidationStateCompleted ValidationState = "COMPLETED"
	ValidationStateRejected  ValidationState = "REJECTED"
	ValidationStateFailed    ValidationState = "FAILED"
)

// Verdict describes the final public result of a validation request.
type Verdict string

const (
	VerdictVerified     Verdict = "VERIFIED"
	VerdictRejected     Verdict = "REJECTED"
	VerdictInconclusive Verdict = "INCONCLUSIVE"
)

// ReasonCode provides a stable public classification for a verdict.
//
// Reason codes are intentionally implementation-independent and must not
// expose Protected Runtime decision inputs or internal execution logic.
type ReasonCode string

const (
	ReasonContractPreserved      ReasonCode = "CONTRACT_PRESERVED"
	ReasonEvidenceVerified       ReasonCode = "EVIDENCE_VERIFIED"
	ReasonCompatibilityConfirmed ReasonCode = "COMPATIBILITY_CONFIRMED"
	ReasonInvalidRequest         ReasonCode = "INVALID_REQUEST"
	ReasonIntegrityFailure       ReasonCode = "INTEGRITY_FAILURE"
	ReasonContractViolation      ReasonCode = "CONTRACT_VIOLATION"
	ReasonInsufficientEvidence   ReasonCode = "INSUFFICIENT_EVIDENCE"
)

// EvidenceRetrievalStatus describes whether an evidence artifact may be
// retrieved through the public boundary.
type EvidenceRetrievalStatus string

const (
	EvidenceAvailable  EvidenceRetrievalStatus = "AVAILABLE"
	EvidenceRestricted EvidenceRetrievalStatus = "RESTRICTED"
	EvidenceExpired    EvidenceRetrievalStatus = "EXPIRED"
)

// PublicHealth contains bounded public service-health information.
type PublicHealth struct {
	Schema          string       `json:"schema"`
	Status          HealthStatus `json:"status"`
	ObservedAt      time.Time    `json:"observed_at"`
	BoundaryVersion string       `json:"boundary_version"`
}

// PublicCapabilities declares public integration capabilities supported by
// an authorized endpoint.
type PublicCapabilities struct {
	Schema          string   `json:"schema"`
	BoundaryVersion string   `json:"boundary_version"`
	Capabilities    []string `json:"capabilities"`
}

// CreateValidationRequest represents a public validation request.
//
// Metadata must contain only non-sensitive, public values. Credentials,
// secrets, customer infrastructure details, and Protected Runtime state must
// never be submitted through this structure.
type CreateValidationRequest struct {
	Schema                string         `json:"schema"`
	ClientRequestID       string         `json:"client_request_id"`
	Scenario              string         `json:"scenario"`
	RequestedCapabilities []string       `json:"requested_capabilities,omitempty"`
	Metadata              map[string]any `json:"metadata,omitempty"`
}

// ValidationRequestAccepted confirms that a public validation request was
// admitted for processing.
type ValidationRequestAccepted struct {
	Schema     string          `json:"schema"`
	RequestID  string          `json:"request_id"`
	Status     ValidationState `json:"status"`
	AcceptedAt time.Time       `json:"accepted_at"`
}

// ValidationStatus contains the externally observable state of a validation
// request.
type ValidationStatus struct {
	Schema           string          `json:"schema"`
	RequestID        string          `json:"request_id"`
	Status           ValidationState `json:"status"`
	UpdatedAt        time.Time       `json:"updated_at"`
	VerdictAvailable bool            `json:"verdict_available"`
}

// EvidenceReference identifies an externally verifiable evidence artifact.
type EvidenceReference struct {
	Schema          string                  `json:"schema"`
	EvidenceID      string                  `json:"evidence_id"`
	Digest          string                  `json:"digest"`
	CreatedAt       time.Time               `json:"created_at"`
	MediaType       string                  `json:"media_type,omitempty"`
	RetrievalStatus EvidenceRetrievalStatus `json:"retrieval_status"`
	PublicMetadata  map[string]any          `json:"public_metadata,omitempty"`
}

// PublicVerdict represents a final implementation-independent validation
// result.
type PublicVerdict struct {
	Schema      string             `json:"schema"`
	RequestID   string             `json:"request_id"`
	Verdict     Verdict            `json:"verdict"`
	ReasonCode  ReasonCode         `json:"reason_code"`
	CompletedAt time.Time          `json:"completed_at"`
	Evidence    *EvidenceReference `json:"evidence,omitempty"`
}