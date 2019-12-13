package test_test

import (
	"testing"

	"github.com/alexeyco/unisender/test"
)

func TestEqualFloat64(t *testing.T) {
	if !test.EqualFloat64(1.2, 1.2) {
		t.Error(`Should be equal`)
	}

	if test.EqualFloat64(1.2, 1.21) {
		t.Error(`Error should not be equal`)
	}
}
