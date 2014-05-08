//
// Tests for CsvFileReader
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package csv_reader

import (
	"os"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	csvfile, _ := os.Open("./example/charity.csv")
	r := NewCsvFileReader(csvfile)
	strings, err, _ := r.GetStrings("Code", "Text Deutsch")
	if err != nil {
		t.Errorf(err.Error())
	}
	key := "about.heavystuff.headline"
	expected := "Schwere Kost"
	actual := strings["about"].(map[string]interface{})["heavystuff"].(map[string]interface{})["headline"]
	if actual != expected {
		t.Errorf("Strings not parsed. Got '%s' for key '%s' instead of '%s'!", actual, key, expected)
	}
}

func TestInvalidKey(t *testing.T) {
	r := NewCsvFileReader(strings.NewReader("key,de\ninvalid key,Wert"))
	_, err, dataErrors := r.GetStrings("key", "de")
	assertNotNil(t, err, "GetStrings error")
	assertEquals(t, len(dataErrors), 1)
	assertEquals(t, dataErrors[0].Key, "invalid key")
	assertEquals(t, dataErrors[0].Msg, "The key \"invalid key\" is invalid. Please only use A-z, 0-9, dots for seperation and no numbers in the beginning or after dots.")
}

func TestDuplicateKey(t *testing.T) {
	r := NewCsvFileReader(strings.NewReader("key,de\nthe.key,Wert\nthe.key,Wert2"))
	_, err, dataErrors := r.GetStrings("key", "de")
	assertNotNil(t, err, "GetStrings error")
	assertEquals(t, len(dataErrors), 1)
	assertEquals(t, dataErrors[0].Key, "the.key")
	assertEquals(t, dataErrors[0].Msg, "The key \"the.key\" is used twice.")
}

func TestOverlappingKey(t *testing.T) {
	r := NewCsvFileReader(strings.NewReader("key,de\nthe.key.overlaps,Wert\nthe.key.overlaps.this,Wert2"))
	_, err, dataErrors := r.GetStrings("key", "de")
	assertNotNil(t, err, "GetStrings error")
	assertEquals(t, len(dataErrors), 1)
	assertEquals(t, dataErrors[0].Key, "the.key.overlaps.this")
	assertEquals(t, dataErrors[0].Msg, "A prefix of the key \"the.key.overlaps.this\" is already in use.")
}

func assertEquals(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("Value '%s' does not match '%s'.", actual, expected)
	}
}

func assertNotNil(t *testing.T, item interface{}, what string) {
	if item == nil {
		t.Errorf("Failed asserting that %s is not nil", what)
	}
}
