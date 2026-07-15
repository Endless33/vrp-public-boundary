# VRP Public Engineering Boundary v0.1.0

## Release Overview

Version 0.1.0 establishes the first public engineering boundary of the VRP ecosystem.

This release provides implementation-independent documentation, contracts, schemas, integration examples, validation concepts, and evidence models without exposing the Protected Runtime.

---

## Included

- Public engineering documentation
- Runtime boundary definitions
- Open Engineering Model
- OpenAPI specification
- Protocol Buffers contracts
- JSON schemas
- Go integration examples
- Python integration examples
- HTTP request examples
- Public sample payloads
- Validation documentation
- Evidence verification examples
- Pilot documentation
- Compatibility and versioning policies
- Engineering governance documentation

---

## Explicitly Excluded

This release does not contain:

- Protected Runtime
- Runtime decision engine
- Authority implementation
- Recovery implementation
- Session continuity internals
- Transport migration internals
- Replay protection implementation
- Proprietary mathematical models
- Internal state machines
- Customer-specific runtime profiles
- Commercial deployment packages

---

## Release Purpose

The purpose of this release is to allow external engineers to:

- understand the public engineering surface;
- inspect versioned public contracts;
- review validation and evidence models;
- prepare technical evaluations;
- plan future Pilot integrations.

The release does not provide a runnable production runtime.

---

## Endpoint Status

All endpoint addresses used in examples are illustrative placeholders.

No production public endpoint is declared by this repository.

Authorized Pilot environments may receive dedicated endpoints, credentials, runtime packages, and deployment profiles through separate commercial agreements.

---

## Stability

Current maturity:

```text
PUBLIC_PREVIEW
```

Public contracts and schemas may evolve before the first stable release.

Breaking public changes will always be introduced through explicit versioning.

---

## Planned Next Steps

- Initial Go SDK
- Initial Python SDK
- Expanded Protocol Buffers
- Additional JSON Schemas
- More integration examples
- Public contract validation tools
- Pilot preparation materials
- VRP Pilot Kit repository

---

## Engineering Principle

Public engineering interfaces should remain transparent.

Protected implementation should remain proprietary.

Independent validation should remain reproducible.

---

## Release Summary

```text
RELEASE=VRP_PUBLIC_ENGINEERING_BOUNDARY_V0.1.0
PUBLIC_CONTRACTS=AVAILABLE
PUBLIC_SCHEMAS=AVAILABLE
PUBLIC_API=AVAILABLE
OPENAPI=AVAILABLE
PROTOBUF=AVAILABLE
HTTP_EXAMPLES=AVAILABLE
SDK_EXAMPLES=AVAILABLE
INTEGRATION_DOCUMENTATION=AVAILABLE
PILOT_DOCUMENTATION=AVAILABLE
PROTECTED_RUNTIME_INCLUDED=NO
STATUS=PUBLIC_PREVIEW
```

Engineering continues.
```