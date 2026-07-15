# Non-Goals

## Purpose

This document defines what the VRP Public Engineering Boundary is not intended to provide.

Clear non-goals prevent incorrect expectations, architectural confusion, and accidental expansion of the public disclosure boundary.

## Not a Complete Runtime

This repository is not a complete implementation of VRP.

It cannot be used to build, reproduce, or replace the Protected Runtime.

## Not an Open Runtime Release

The existence of public contracts, schemas, SDK surfaces, and examples does not mean that the runtime implementation is open source.

The public engineering boundary and the Protected Runtime are separate artifacts.

## Not a Public Production Service

Example endpoints are placeholders.

This repository does not declare a public production API, hosted validation service, or unrestricted runtime endpoint.

## Not an Anonymous Trial

The VRP Pilot Program is not an anonymous public beta.

Runtime access, credentials, deployment packages, and customer-specific profiles are issued only through authorized evaluation or commercial processes.

## Not an Implementation Specification

Public documents describe:

- observable behavior;
- integration boundaries;
- public data structures;
- validation workflows;
- evidence formats.

They do not describe:

- internal execution order;
- protected state transitions;
- decision algorithms;
- authority selection mechanisms;
- recovery calculations;
- proprietary mathematics.

## Not a Reconstruction Guide

No artifact in this repository should materially assist reconstruction of the Protected Runtime.

Information that would weaken this boundary remains private.

## Not a Replacement Claim

This repository does not claim that VRP automatically replaces every existing networking protocol, transport, VPN product, operating-system network stack, or security layer.

VRP is evaluated according to defined use cases, integration requirements, and deployment constraints.

## Not a Security Guarantee by Documentation Alone

Public schemas and contracts do not independently guarantee production security.

Security depends on:

- correct deployment;
- authorized runtime components;
- configuration;
- operational controls;
- threat model;
- environment;
- ongoing validation.

## Not a Free Commercial License

Publication does not grant unrestricted rights to:

- commercialize VRP;
- redistribute protected components;
- operate commercial services using protected technology;
- create derivative commercial runtimes;
- use VRP trademarks without permission.

Applicable rights are defined by the repository license and separate written agreements.

## Not a Substitute for Pilot Validation

Documentation provides an engineering foundation.

Production suitability must be established through environment-specific evaluation, validation, and deployment review.

## Core Rule

The repository explains how external systems may interact with VRP.

It does not explain how to recreate VRP.

## Status

```text
DOCUMENT_STATUS=ACTIVE
PUBLIC_BOUNDARY=ENFORCED
PROTECTED_RUNTIME_INCLUDED=NO