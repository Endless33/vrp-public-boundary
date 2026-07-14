# Public API Policy

## Purpose

Public APIs define stable interaction points between external systems and the VRP engineering ecosystem.

Public APIs are designed for interoperability.

They are not intended to expose runtime implementation.

---

## Design Principles

Every public API should be:

- deterministic
- documented
- versioned
- backward conscious
- implementation independent

---

## Public Responsibilities

Public APIs may provide:

- capability discovery
- validation requests
- evidence retrieval
- compatibility information
- public runtime metadata

---

## Protected Responsibilities

Public APIs never expose:

- runtime execution
- internal state
- authority transitions
- recovery algorithms
- proprietary mathematics
- protected runtime logic

---

## Compatibility

Breaking API changes should be introduced through explicit versioning.

Existing integrations should remain stable whenever practical.

---

STATUS: ACTIVE