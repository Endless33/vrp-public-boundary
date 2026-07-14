package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type EvidenceReference struct {
	Schema          string         `json:"schema"`
	EvidenceID      string         `json:"evidence_id"`
	Digest          string         `json:"digest"`
	CreatedAt       time.Time      `json:"created_at"`
	MediaType       string         `json:"media_type,omitempty"`
	RetrievalStatus string         `json:"retrieval_status"`
	PublicMetadata  map[string]any `json:"public_metadata,omitempty"`
}

func loadReference(path string) (EvidenceReference, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return EvidenceReference{}, fmt.Errorf("read reference: %w", err)
	}

	var reference EvidenceReference

	if err := json.Unmarshal(data, &reference); err != nil {
		return EvidenceReference{}, fmt.Errorf("decode reference: %w", err)
	}

	if reference.Schema != "vrp.evidence-reference/v1" {
		return EvidenceReference{}, fmt.Errorf(
			"unsupported schema: %q",
			reference.Schema,
		)
	}

	if reference.EvidenceID == "" {
		return EvidenceReference{}, errors.New("missing evidence_id")
	}

	if reference.RetrievalStatus == "" {
		return EvidenceReference{}, errors.New(
			"missing retrieval_status",
		)
	}

	return reference, nil
}

func calculateDigest(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read evidence: %w", err)
	}

	sum := sha256.Sum256(data)

	return "sha256:" + hex.EncodeToString(sum[:]), nil
}

func normalizeDigest(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(
			os.Stderr,
			"usage: %s <reference.json> <evidence-file>\n",
			os.Args[0],
		)
		os.Exit(2)
	}

	referencePath := os.Args[1]
	evidencePath := os.Args[2]

	reference, err := loadReference(referencePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "REFERENCE_INVALID=%v\n", err)
		os.Exit(1)
	}

	actualDigest, err := calculateDigest(evidencePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "EVIDENCE_READ_FAILED=%v\n", err)
		os.Exit(1)
	}

	expectedDigest := normalizeDigest(reference.Digest)
	actualDigest = normalizeDigest(actualDigest)

	fmt.Printf("SCHEMA=%s\n", reference.Schema)
	fmt.Printf("EVIDENCE_ID=%s\n", reference.EvidenceID)
	fmt.Printf("EXPECTED_DIGEST=%s\n", expectedDigest)
	fmt.Printf("ACTUAL_DIGEST=%s\n", actualDigest)

	if expectedDigest != actualDigest {
		fmt.Println("FINAL_VERDICT=EVIDENCE_INTEGRITY_REJECTED")
		os.Exit(1)
	}

	fmt.Println("FINAL_VERDICT=EVIDENCE_INTEGRITY_VERIFIED")
}