//
// Tests for YamlLangWriter
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package yaml

import (
	"bytes"
	assert "github.com/dothiv/translations-updater/testing"
	"strings"
	"testing"
)

// Test for YamlLangWriter
func TestWriteFile(t *testing.T) {
	str := make(map[string]interface{})
	str["test"] = make(map[string]string)
	str["test"].(map[string]string)["key"] = "value"
	str["another"] = make(map[string]interface{})
	str["another"].(map[string]interface{})["test"] = make(map[string]int)
	str["another"].(map[string]interface{})["test"].(map[string]int)["key"] = 42

	w := NewYamlLangWriter()
	target := bytes.NewBufferString("")
	w.WriteTo(str, target)
	yaml := target.String()
	expected := "another:\n  test:\n    key: 42\ntest:\n  key: value"
	assert.Equals(t, expected, strings.TrimSpace(yaml))
}
