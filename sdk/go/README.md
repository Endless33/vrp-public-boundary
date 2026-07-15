# Go SDK

## Overview

The Go SDK provides a lightweight client for interacting with the VRP Public Engineering Boundary.

The SDK is intended for engineers, system architects, validation partners, and organizations preparing Pilot integrations.

---

## Current Features

- Public Health
- Capability Discovery
- Validation Status
- Public Verdict Retrieval
- Evidence Reference Retrieval
- Stable Public Types
- Public Error Model

---

## Not Included

The Go SDK does **not** contain:

- Protected Runtime
- Runtime Algorithms
- Authority Logic
- Recovery Logic
- Replay Logic
- Runtime Mathematics
- Commercial Runtime Components

---

## Installation

```bash
go get github.com/Endless33/vrp-public-boundary/sdk/go
```

---

## Example

```go
client := vrpboundary.NewClient(
    "https://pilot.example.invalid/api/v1",
)

health, err := client.GetHealth(context.Background())
if err != nil {
    panic(err)
}

fmt.Println(health.Status)
```

---

## Version

Current SDK Version

```text
0.1.0
```

---

Engineering continues.