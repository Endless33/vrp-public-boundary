"""
VRP Public Engineering Boundary SDK.

This package provides a lightweight Python interface for the
public VRP engineering boundary.

The SDK intentionally exposes only public engineering interfaces.

Protected Runtime components are not included.
"""

from .client import VRPPublicClient
from .models import (
    PublicHealth,
    PublicCapabilities,
    ValidationStatus,
    PublicVerdict,
    EvidenceReference,
)

__version__ = "0.1.0"

__all__ = [
    "VRPPublicClient",
    "PublicHealth",
    "PublicCapabilities",
    "ValidationStatus",
    "PublicVerdict",
    "EvidenceReference",
]