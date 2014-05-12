//
// Tests for CsvFileReader
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package csv

import (
	assert "github.com/dothiv/translations-updater/testing"
	"os"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	csvfile, _ := os.Open("../example/charity.csv")
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
	assert.NotNil(t, err, "GetStrings error")
	assert.Equals(t, len(dataErrors), 1)
	assert.Equals(t, dataErrors[0].Key, "invalid key")
	assert.Equals(t, dataErrors[0].Msg, "The key \"invalid key\" is invalid. Please only use A-z, 0-9, dots for seperation and no numbers in the beginning or after dots.")
}

func TestDuplicateKey(t *testing.T) {
	r := NewCsvFileReader(strings.NewReader("key,de\nthe.key,Wert\nthe.key,Wert2"))
	_, err, dataErrors := r.GetStrings("key", "de")
	assert.NotNil(t, err, "GetStrings error")
	assert.Equals(t, len(dataErrors), 1)
	assert.Equals(t, dataErrors[0].Key, "the.key")
	assert.Equals(t, dataErrors[0].Msg, "The key \"the.key\" is used twice.")
}

func TestOverlappingKey(t *testing.T) {
	r := NewCsvFileReader(strings.NewReader("key,de\nthe.key.overlaps,Wert\nthe.key.overlaps.this,Wert2"))
	_, err, dataErrors := r.GetStrings("key", "de")
	assert.NotNil(t, err, "GetStrings error")
	assert.Equals(t, len(dataErrors), 1)
	assert.Equals(t, dataErrors[0].Key, "the.key.overlaps.this")
	assert.Equals(t, dataErrors[0].Msg, "A prefix of the key \"the.key.overlaps.this\" is already in use.")
}

func TestIgnoreEmptyKey(t *testing.T) {
	r := NewCsvFileReader(strings.NewReader("key,de\nhas.key,Wert\n,Wert2"))
	_, err, _ := r.GetStrings("key", "de")
	assert.Nil(t, err, "GetStrings error")
}
