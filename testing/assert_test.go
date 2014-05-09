//
// Tests for assertions
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package testing

import (
	"testing"
)

func TestEquals(t *testing.T) {
	mockTest := new(testing.T)
	Equals(mockTest, "a", "a")
	if mockTest.Failed() {
		t.Errorf("It should not fail.")
	}
	Equals(mockTest, "a", "b")
	if !mockTest.Failed() {
		t.Errorf("It should fail.")
	}
}

func TestNotNil(t *testing.T) {
	mockTest := new(testing.T)
	NotNil(mockTest, "a", "a")
	if mockTest.Failed() {
		t.Errorf("It should not fail.")
	}
	NotNil(mockTest, nil, "a")
	if !mockTest.Failed() {
		t.Errorf("It should fail.")
	}
}

func TestNil(t *testing.T) {
	mockTest := new(testing.T)
	Nil(mockTest, nil, "a")
	if mockTest.Failed() {
		t.Errorf("It should not fail.")
	}
	Nil(mockTest, "a", "a")
	if !mockTest.Failed() {
		t.Errorf("It should fail.")
	}
}
