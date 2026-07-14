# VRP Engineering Principles

## Core Principles

The public engineering boundary of VRP follows several immutable principles.

---

## Determinism

Observable behavior should remain deterministic under identical inputs.

---

## Reproducibility

Validation results should be reproducible by independent parties.

---

## Continuity First

Session continuity takes priority over transport continuity.

---

## Observable Verification

External systems should verify exported evidence without requiring access to the protected runtime.

---

## Stable Contracts

Public interfaces should evolve through explicit versioning.

Breaking changes should be intentional.

---

## Implementation Independence

Public engineering contracts describe observable behavior.

They never describe the internal implementation.

---

## Security by Boundary

Public documentation should never disclose protected runtime logic.

Knowledge of public contracts must not reveal proprietary implementation.

---

## Engineering Philosophy

Interfaces may be public.

Evidence may be public.

Verification may be public.

Protected Runtime remains protected.

---

STATUS: ACTIVE