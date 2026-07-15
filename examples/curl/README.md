# cURL Examples

## Overview

The examples below demonstrate how a client may interact with the VRP Public Engineering Boundary.

These requests communicate only with public engineering interfaces.

They never expose Protected Runtime functionality.

---

## Health

```bash
curl \
  -X GET \
  "https://pilot.example.invalid/api/v1/health" \
  -H "Accept: application/json"
```

---

## Capabilities

```bash
curl \
  -X GET \
  "https://pilot.example.invalid/api/v1/capabilities" \
  -H "Accept: application/json"
```

---

## Validation Status

```bash
curl \
  -X GET \
  "https://pilot.example.invalid/api/v1/validation/<request-id>/status" \
  -H "Accept: application/json"
```

---

## Validation Verdict

```bash
curl \
  -X GET \
  "https://pilot.example.invalid/api/v1/validation/<request-id>/verdict" \
  -H "Accept: application/json"
```

---

## Evidence Reference

```bash
curl \
  -X GET \
  "https://pilot.example.invalid/api/v1/evidence/<evidence-id>" \
  -H "Accept: application/json"
```

---

## Notes

Example endpoints are placeholders.

Production endpoints are provided only to authorized Pilot environments.

Protected Runtime functionality is intentionally excluded from the public engineering boundary.

Engineering continues.