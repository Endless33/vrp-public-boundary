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

	caps, err := client.GetCapabilities(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("======= PUBLIC CAPABILITIES =======")

	for _, c := range caps.Capabilities {
		fmt.Println("-", c)
	}
}