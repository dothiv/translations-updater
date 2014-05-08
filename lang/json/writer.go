//
// LangWriter
//
// Used to write our JSON lang files.
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package lang

import (
	"encoding/json"
	"io"
)

// This is the main object
type JsonLangWriter struct {
}

func NewJsonLangWriter() (w *JsonLangWriter) {
	w = new(JsonLangWriter)
	return
}

// Write strings to the target
func (l *JsonLangWriter) WriteTo(str map[string]interface{}, target io.Writer) (err error) {
	encoder := json.NewEncoder(target)
	err = encoder.Encode(str)
	return
}
