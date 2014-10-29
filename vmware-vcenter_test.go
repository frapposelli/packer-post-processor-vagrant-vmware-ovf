package main

import (
	"testing"
)

func TestVMwarevCenterProvider_impl(t *testing.T) {
	var _ Provider = new(VMwarevCenterProvider)
}
