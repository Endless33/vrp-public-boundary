# Security Model

## Engineering Philosophy

VRP separates public engineering interfaces from protected runtime implementation.

This separation is intentional.

---

## Public Layer

The public layer contains:

- documentation
- contracts
- schemas
- SDK interfaces
- evidence formats
- validation specifications

---

## Protected Layer

The protected layer contains:

- runtime implementation
- proprietary algorithms
- mathematical models
- authority logic
- recovery mechanisms
- migration engine
- security internals

---

## Security Principle

Knowledge of the public layer must not enable reconstruction of the protected runtime.

---

STATUS: ACTIVE