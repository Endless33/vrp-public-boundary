# Security Policy

## Purpose

This document describes how security-related issues concerning the VRP Public Engineering Boundary should be reported and handled.

The purpose is to improve responsible engineering practices while protecting both reporters and users.

---

## Scope

This policy applies only to the public engineering artifacts contained in this repository.

Examples include:

- documentation
- OpenAPI specifications
- Protocol Buffers contracts
- JSON Schemas
- SDK examples
- HTTP examples
- public integration guides
- engineering documentation

---

## Out of Scope

This repository does not contain:

- Protected Runtime
- commercial runtime packages
- runtime decision engine
- authority engine
- recovery engine
- migration engine
- replay engine
- runtime mathematics
- customer deployments

Security reports concerning components that are not distributed through this repository cannot be evaluated through GitHub.

---

## Reporting a Security Issue

Please do not publish active vulnerabilities before responsible disclosure.

Security reports should include:

- affected artifact
- repository version
- reproduction steps
- expected behavior
- observed behavior
- supporting logs when appropriate

Reports may be submitted to:

jumpingvpn@proton.me

---

## Responsible Disclosure

Please allow reasonable time for investigation before public disclosure.

If the reported issue affects public engineering artifacts, updates will be published through normal repository releases.

---

## Protected Runtime

Questions requesting disclosure of:

- proprietary implementation
- runtime algorithms
- protected mathematics
- authority logic
- recovery logic
- internal architecture

cannot be answered through the security reporting process.

---

## Engineering Principle

Security should improve transparency where appropriate.

Security should never require disclosure of proprietary implementation.

---

## Status

```text
SECURITY_POLICY=ACTIVE
RESPONSIBLE_DISCLOSURE=SUPPORTED
PROTECTED_RUNTIME=PRIVATE
```