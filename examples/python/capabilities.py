from sdk.python.client import VRPPublicClient

client = VRPPublicClient(
    "https://pilot.example.invalid/api/v1"
)

caps = client.get_capabilities()

print("======= PUBLIC CAPABILITIES =======")

for capability in caps.capabilities:
    print("-", capability)