package main

import (
	"testing"
)

func TestVMwareOVFProvider_impl(t *testing.T) {
	var _ Provider = new(VMwareOVFProvider)
}
