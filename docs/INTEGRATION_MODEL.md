# VRP Integration Model

## Purpose

This document describes how external systems interact with the public VRP engineering surface.

It intentionally documents only observable integration behavior.

Protected Runtime implementation is not included.

---

## High-Level Model

```text
External System
        │
        ▼
Public SDK / Contract
        │
        ▼
Protected Runtime
        │
        ▼
Evidence / Verdict
```

The internal runtime architecture is intentionally omitted.

---

## Public Responsibilities

External systems may:

- submit requests
- receive public verdicts
- verify exported evidence
- integrate through public contracts
- consume future SDK interfaces

---

## Protected Responsibilities

The Protected Runtime is responsible for:

- runtime continuity
- authority management
- recovery logic
- replay protection
- migration logic
- deterministic decisions
- internal state protection

Implementation details remain private.

---

## Design Principles

The public integration layer follows these principles:

- minimal
- deterministic
- versioned
- implementation-independent
- reproducible
- externally verifiable

---

## Pilot Integration

Commercial Pilot deployments may include additional licensed components.

Those components are not part of this repository.

---

## Status

STATUS: DRAFT

Protected Runtime: NOT INCLUDED