//
// YamlLangWriter
//
// Used to write our Yaml lang files.
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package yaml

import (
	"bufio"
	yaml "gopkg.in/yaml.v1"
	"io"
)

const NUM_INDENT = 4

// This is the main object
type YamlLangWriter struct {
}

func NewYamlLangWriter() (w *YamlLangWriter) {
	w = new(YamlLangWriter)
	return
}

// Write strings to the target
func (l *YamlLangWriter) WriteTo(str map[string]interface{}, target io.Writer) (err error) {
	w := bufio.NewWriter(target)

	d := []byte{}

	d, err = yaml.Marshal(&str)
	if err != nil {
		return
	}
	w.Write(d)
	w.Flush()
	return
}
