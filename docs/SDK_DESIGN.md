# SDK Design

## Purpose

The VRP SDK provides a stable engineering interface between external applications and the VRP ecosystem.

The SDK is intended to simplify integration while preserving complete separation from the Protected Runtime.

---

## Design Objectives

The SDK is designed to be:

- deterministic
- lightweight
- versioned
- implementation independent
- easy to integrate
- easy to validate

---

## Engineering Principles

The SDK should expose engineering interfaces.

It should never expose runtime implementation.

Applications interact with documented contracts rather than proprietary runtime internals.

---

## Public Responsibilities

Future SDK components may include:

- client libraries
- request builders
- validation helpers
- evidence utilities
- compatibility helpers
- integration examples

---

## Protected Responsibilities

The SDK never contains:

- runtime execution
- authority algorithms
- recovery engine
- migration engine
- replay implementation
- runtime mathematics

---

## Evolution

SDK interfaces will evolve independently from the Protected Runtime.

Public compatibility will be maintained whenever practical.

---

STATUS: ACTIVE