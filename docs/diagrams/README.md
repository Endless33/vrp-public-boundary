# Architecture Diagrams

## Purpose

This directory contains engineering diagrams describing the public VRP engineering boundary.

The diagrams are intended to help engineers understand public interfaces and integration workflows without exposing Protected Runtime implementation.

---

## Planned Diagrams

### Public Boundary

Illustrates the separation between:

- External Systems
- Public Engineering Boundary
- Protected Runtime

---

### Validation Flow

Illustrates the public validation lifecycle.

Client

↓

Validation Request

↓

Validation Processing

↓

Public Verdict

↓

Evidence Reference

---

### Evidence Flow

Illustrates how evidence is referenced, verified and consumed by external systems.

---

### Pilot Workflow

Illustrates the engineering workflow for Pilot onboarding.

Application

↓

Technical Review

↓

Pilot Approval

↓

Environment Preparation

↓

Validation

↓

Engineering Feedback

---

### SDK Architecture

Illustrates the relationship between:

Go SDK

↓

Public API

↓

Engineering Boundary

↓

Protected Runtime

---

### Public API

Illustrates externally observable endpoints.

Health

Capabilities

Validation

Evidence

Verdicts

---

## File Formats

Future releases may include:

- SVG
- Draw.io
- PNG
- PDF

---

## Boundary Rule

Architecture diagrams describe public engineering behavior only.

Protected Runtime implementation is intentionally excluded.

---

Engineering continues.