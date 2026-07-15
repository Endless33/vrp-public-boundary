from sdk.python.client import VRPPublicClient

client = VRPPublicClient(
    "https://pilot.example.invalid/api/v1"
)

health = client.get_health()

print("========== VRP PUBLIC HEALTH ==========")
print("Schema:", health.schema)
print("Status:", health.status)
print("Boundary:", health.boundary_version)
print("Observed:", health.observed_at)