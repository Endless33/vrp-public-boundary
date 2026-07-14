# Public API

## Overview

This directory is reserved for public API specifications exposed by the VRP engineering boundary.

The APIs described here are intended for integration, validation and pilot environments.

Protected Runtime implementation is intentionally excluded.

---

## Planned APIs

Future public APIs may include:

- runtime discovery
- capability negotiation
- evidence export
- validation requests
- verdict retrieval
- health information
- compatibility checks

---

## Design Goals

Every public API should be:

- deterministic
- versioned
- implementation independent
- backward conscious
- externally verifiable

---

## Excluded

This repository will never expose:

- runtime internals
- authority implementation
- recovery implementation
- mathematical models
- state machines
- protected algorithms

---

STATUS: RESERVED