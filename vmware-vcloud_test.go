package main

import (
	"testing"
)

func TestVMwarevCloudProvider_impl(t *testing.T) {
	var _ Provider = new(VMwarevCloudProvider)
}
