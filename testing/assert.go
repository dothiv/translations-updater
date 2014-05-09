//
// Simple assertions for testing
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package testing

import (
	"testing"
)

func Equals(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("Value '%s' does not match '%s'.", actual, expected)
	}
}

func NotNil(t *testing.T, item interface{}, what string) {
	if item == nil {
		t.Errorf("Failed asserting that %s is not nil", what)
	}
}

func Nil(t *testing.T, item interface{}, what string) {
	if item != nil {
		t.Errorf("Failed asserting that %s is nil", what)
	}
}
