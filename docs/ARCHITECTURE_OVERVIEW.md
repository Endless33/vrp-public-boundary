# Architecture Overview

## Overview

VRP is designed around a layered engineering architecture.

This repository documents the public engineering boundary only.

The Protected Runtime is intentionally excluded.

---

## High-Level Architecture

```text
Applications
        │
        ▼
Public SDK
        │
        ▼
Public Contracts
        │
        ▼
Protected Runtime Boundary
        │
        ▼
Protected Runtime
        │
        ▼
Evidence Export
```

The diagram illustrates engineering responsibilities only.

It does not describe runtime implementation.

---

## Public Engineering Layer

The public layer is responsible for:

- documentation
- SDK interfaces
- contracts
- schemas
- validation
- evidence
- integration guidance

This layer is implementation-independent.

---

## Protected Runtime Layer

The protected layer contains:

- runtime execution
- authority management
- recovery mechanisms
- transport migration
- replay protection
- deterministic decision engine
- internal protocol mathematics

These components remain proprietary.

---

## Separation Principle

The public engineering boundary allows engineers to:

- integrate
- validate
- evaluate
- prepare Pilot deployments

without requiring access to runtime implementation.

---

## Long-Term Direction

Future releases will expand the engineering boundary through:

- SDK evolution
- additional contracts
- validation specifications
- Pilot documentation
- public schemas
- integration examples

The Protected Runtime will continue to evolve independently.

---

STATUS: ACTIVE