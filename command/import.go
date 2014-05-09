//
// Import strings command
//
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package command

import (
	"github.com/dothiv/translations-updater/csv"
	json "github.com/dothiv/translations-updater/lang/json"
	"github.com/dothiv/translations-updater/util"
	"os"
)

type ImportCommand struct {
	source  string
	target  string
	codeCol string
	valCol  string
}

func NewImportCommand(src string, target string, codeCol string, valCol string) (c *ImportCommand) {
	c = new(ImportCommand)
	c.source = src
	c.target = target
	c.codeCol = codeCol
	c.valCol = valCol
	return
}

func (c *ImportCommand) Exec() (err error, errorStrings []csv.KeyError) {
	var csvfile *os.File
	csvfile, err = util.LoadUri(c.source)
	if err != nil {
		return
	}

	r := csv.NewCsvFileReader(csvfile)

	var str map[string]interface{}

	str, err, errorStrings = r.GetStrings(c.codeCol, c.valCol)
	if err != nil {
		return
	}

	var jsfile *os.File
	jsfile, err = os.OpenFile(c.target, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	defer jsfile.Close()

	w := json.NewJsonIndentLangWriter()
	w.WriteTo(str, jsfile)
	return
}
