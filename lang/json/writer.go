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
	"bufio"
	"encoding/json"
	"io"
)

// This is the main object
type JsonLangWriter struct {
	indent bool
}

func NewJsonLangWriter() (w *JsonLangWriter) {
	w = new(JsonLangWriter)
	return
}

// Write strings to the target
func (l *JsonLangWriter) WriteTo(str map[string]interface{}, target io.Writer) (err error) {
	if l.indent {
		w := bufio.NewWriter(target)
		b, _ := json.MarshalIndent(str, "", "    ")
		w.Write(b)
	} else {
		encoder := json.NewEncoder(target)
		err = encoder.Encode(str)
	}
	return
}
