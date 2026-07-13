# VRP Public Engineering Boundary

## Overview

This repository contains the public engineering boundary of the VRP (Veil Routing Protocol) ecosystem.

It is intended for engineers, system architects, validation partners, and pilot participants who require a documented and reproducible integration surface without exposing the protected runtime implementation.

---

## Repository Scope

Included:

- Engineering contracts
- Public SDK interfaces
- Validation specifications
- Evidence formats
- Integration documentation
- Public APIs
- Example client implementations
- Pilot integration guides

---

## Not Included

This repository does **not** contain:

- Protected Runtime
- Runtime decision engine
- Authority algorithms
- Recovery engine
- Session migration implementation
- Internal mathematics
- Proprietary runtime logic

These components remain protected.

---

## Purpose

The objective is to provide a stable engineering interface that allows independent validation, integration, and technical evaluation without exposing protected implementation details.

The engineering boundary is public.

The runtime core is not.

---

## Repository Structure

```
docs/
sdk/
api/
contracts/
examples/
validation/
evidence/
schemas/
```

Additional components will be introduced as the engineering boundary expands.

---

## Engineering Principles

- Reproducible
- Verifiable
- Deterministic
- Vendor Neutral
- Continuity First

---

## Status

Engineering Boundary

Initial Public Release

Development Continues.