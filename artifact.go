package main

import (
	"fmt"
	"os"
)

const BuilderId = "frapposelli.post-processor.vagrant-vmware-ovf"

type Artifact struct {
	Path     string
	Provider string
}

func NewArtifact(provider, path string) *Artifact {
	return &Artifact{
		Path:     path,
		Provider: provider,
	}
}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (a *Artifact) Files() []string {
	return []string{a.Path}
}

func (a *Artifact) Id() string {
	return ""
}

func (a *Artifact) String() string {
	return fmt.Sprintf("'%s' provider box: %s", a.Provider, a.Path)
}

func (a *Artifact) Destroy() error {
	return os.Remove(a.Path)
}
