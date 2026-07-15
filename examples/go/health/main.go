package main

import (
	"context"
	"fmt"
	"log"

	vrpboundary "github.com/Endless33/vrp-public-boundary/sdk/go"
)

func main() {
	client := vrpboundary.NewClient(
		"https://pilot.example.invalid/api/v1",
	)

	health, err := client.GetHealth(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("========== VRP PUBLIC HEALTH ==========")
	fmt.Printf("Schema            : %s\n", health.Schema)
	fmt.Printf("Status            : %s\n", health.Status)
	fmt.Printf("Boundary Version  : %s\n", health.BoundaryVersion)
	fmt.Printf("Observed At       : %s\n", health.ObservedAt)
}