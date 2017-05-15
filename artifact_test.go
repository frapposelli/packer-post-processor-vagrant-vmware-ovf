package main

import (
	"github.com/hashicorp/packer/packer"
	"testing"
)

func TestArtifact_ImplementsArtifact(t *testing.T) {
	var raw interface{}
	raw = &Artifact{}
	if _, ok := raw.(packer.Artifact); !ok {
		t.Fatalf("Artifact should be a Artifact")
	}
}

func TestArtifact_Id(t *testing.T) {
	artifact := NewArtifact("vmware_ovf", "./")
	if artifact.Id() != "vmware_ovf" {
		t.Fatalf("should return name as Id")
	}
}
