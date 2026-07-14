# Public Contracts

This directory will contain future public VRP engineering contracts.

Examples include:

- request formats
- response formats
- evidence contracts
- public verdict definitions
- compatibility contracts
- capability declarations

Only public interfaces belong here.

Protected Runtime contracts are intentionally excluded.

---

## Rules

Public contracts must never expose:

- runtime implementation
- internal algorithms
- protected mathematics
- authority logic
- recovery mechanisms
- private state
- security internals

---

## Versioning

Examples:

vrp.contract/v1

vrp.verdict/v1

vrp.evidence/v1

---

STATUS: PREPARATION