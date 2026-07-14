# Runtime Boundary

## Purpose

This document defines the engineering boundary between publicly documented components and the protected VRP runtime.

The boundary exists to enable independent engineering evaluation while preserving proprietary implementation.

---

# Public Boundary

The following components may be publicly documented.

- Documentation
- SDK interfaces
- API specifications
- Engineering contracts
- Public schemas
- Validation workflows
- Evidence specifications
- Integration examples
- Pilot documentation
- Compatibility information

These artifacts describe interaction with VRP.

They do not describe runtime implementation.

---

# Protected Boundary

The following components remain private.

- Runtime Core
- Runtime Scheduler
- Authority Engine
- Recovery Engine
- Migration Engine
- Session Engine
- Replay Engine
- Internal Mathematics
- Runtime Decision Logic
- Internal State Machines
- Runtime Optimizations
- Commercial Runtime Components
- Customer Runtime Profiles

These components are never published through this repository.

---

# Boundary Objectives

The engineering boundary should allow external engineers to:

- understand system integration
- understand public contracts
- perform independent validation
- verify exported evidence
- prepare Pilot deployments

without exposing proprietary runtime technology.

---

# Engineering Rule

Knowledge of every public document contained in this repository must never be sufficient to reconstruct the Protected Runtime.

If a document increases the ability to reconstruct proprietary implementation, it does not belong in this repository.

---

# Public Commitment

Future public releases will expand the engineering boundary.

They will never expose Protected Runtime implementation.

---

STATUS: ACTIVE