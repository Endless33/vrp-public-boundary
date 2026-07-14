# VRP Public Engineering Boundary

## Purpose

This document defines the boundary between the public VRP engineering surface and the protected VRP runtime.

The public boundary exists to support:

- technical evaluation;
- integration planning;
- validation tooling;
- evidence verification;
- pilot preparation;
- interoperability discussions.

It does not expose the implementation responsible for producing runtime decisions.

## Public Surface

The following categories may be documented publicly:

- integration contracts;
- external request and response formats;
- evidence envelope schemas;
- validation scenario definitions;
- public verdict vocabulary;
- SDK-facing interfaces;
- deployment prerequisites;
- pilot workflow documentation;
- redacted examples.

## Protected Surface

The following components are explicitly outside the public boundary:

- Protected Runtime implementation;
- authority selection and transition mechanisms;
- recovery decision logic;
- session continuity internals;
- migration algorithms;
- replay-window implementation;
- commit admission logic;
- internal state machines;
- proprietary mathematical models;
- secret material and key-management procedures;
- anti-tamper mechanisms;
- private telemetry;
- customer-specific runtime profiles.

## Architectural Rule

The public boundary describes:

> what an external system may submit, observe, validate, and verify.

It does not describe:

> how the Protected Runtime reaches its internal decisions.

## Verification Principle

External verification must be possible without possession of the protected implementation.

Observable behavior may be public.

Protected decision machinery remains private.

## Disclosure Rule

No public contribution, issue, example, schema, or document may introduce information that materially assists reconstruction of the Protected Runtime.

When uncertain, the information remains private.

## Status

Boundary state: INITIAL PUBLIC DEFINITION

Protected Runtime status: NOT INCLUDED

Implementation disclosure: PROHIBITED