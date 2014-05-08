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

func TestWriteFile(t *testing.T) {
	w := NewJsonLangWriter()
	str := make(map[string]interface{})
	str["test"] = make(map[string]string)
	str["test"].(map[string]string)["key"] = "value"
	str["another"] = make(map[string]interface{})
	str["another"].(map[string]interface{})["test"] = make(map[string]int)
	str["another"].(map[string]interface{})["test"].(map[string]int)["key"] = 42
	target := bytes.NewBufferString("")
	w.WriteTo(str, target)
	json := target.String()
	ecpected := "{\"another\":{\"test\":{\"key\":42}},\"test\":{\"key\":\"value\"}}"
	assert.Equals(t, ecpected, strings.TrimSpace(json))
}
