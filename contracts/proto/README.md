# VRP Public Protobuf Contracts

## Purpose

This directory contains the official Protocol Buffers (Protobuf) definitions for the VRP Public Engineering Boundary.

These contracts provide a language-neutral interface for integration with the public engineering layer of the VRP ecosystem.

They describe observable interactions only.

They do **not** expose the Protected Runtime.

---

## Goals

The public Protobuf contracts are intended to provide:

- stable engineering interfaces
- cross-language compatibility
- deterministic data structures
- versioned public APIs
- implementation-independent integration

---

## Current Contract

```text
vrp_public_boundary_v1.proto
```

The current contract defines public interfaces for:

- Service Health
- Capability Discovery
- Validation Requests
- Validation Status
- Public Verdict Retrieval
- Evidence Reference Retrieval

Future versions may introduce additional public services while maintaining compatibility whenever practical.

---

## Public Engineering Boundary

The Protobuf contracts describe:

- externally observable requests
- externally observable responses
- public engineering contracts
- public validation workflow
- evidence metadata

The contracts intentionally avoid describing runtime implementation.

---

## Protected Runtime

This directory never contains:

- Runtime Core
- Authority Engine
- Recovery Engine
- Migration Engine
- Session Engine
- Replay Engine
- Internal Mathematics
- Runtime Scheduler
- Runtime Decision Logic
- Customer Runtime Profiles
- Commercial Runtime Components

These components remain proprietary.

---

## Versioning

Current package:

```text
vrp.public.v1
```

Future incompatible public changes will introduce new package versions.

Existing public contracts should remain available during documented compatibility periods whenever practical.

---

## Generated Code

Generated client libraries are intentionally excluded from this repository at this stage.

The `.proto` source files remain the canonical engineering contracts.

Future releases may provide generated SDKs for selected programming languages.

---

## Endpoint Policy

This repository does not publish production service endpoints.

Example endpoints contained in documentation or examples are illustrative only.

Authorized Pilot environments may receive dedicated endpoints through separate commercial agreements.

---

## Engineering Principle

The public contract defines **how** external systems communicate with VRP.

It does not describe **how** the Protected Runtime internally produces its decisions.

Observable interoperability.

Protected implementation.

---

## Status

```text
CONTRACT_STATUS=PUBLIC_PREVIEW
PROTECTED_RUNTIME_INCLUDED=NO
VERSION=v1
ENGINEERING_BOUNDARY=PUBLIC
```