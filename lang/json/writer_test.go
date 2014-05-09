//
// Tests for JsonLangWriter
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package lang

import (
	"bytes"
	assert "github.com/dothiv/translations-updater/testing"
	"strings"
	"testing"
)

// Test for JsonLangWriter
func TestWriteFile(t *testing.T) {
	str := make(map[string]interface{})
	str["test"] = make(map[string]string)
	str["test"].(map[string]string)["key"] = "value"
	str["another"] = make(map[string]interface{})
	str["another"].(map[string]interface{})["test"] = make(map[string]int)
	str["another"].(map[string]interface{})["test"].(map[string]int)["key"] = 42

	w := NewJsonLangWriter()
	target := bytes.NewBufferString("")
	w.WriteTo(str, target)
	json := target.String()
	expected := "{\"another\":{\"test\":{\"key\":42}},\"test\":{\"key\":\"value\"}}"
	assert.Equals(t, expected, strings.TrimSpace(json))
}

// Test for indendet JsonLangWriter
func TestWriteFileIndent(t *testing.T) {
	str := make(map[string]interface{})
	str["a"] = "b"
	expected := "{\n    \"a\": \"b\"\n}"
	w := NewJsonIndentLangWriter()
	target := bytes.NewBufferString("")
	w.WriteTo(str, target)
	json := target.String()
	assert.Equals(t, expected, strings.TrimSpace(json))
}
