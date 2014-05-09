//
// Import strings command
//
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package command

import (
	"github.com/dothiv/translations-updater/config"
	"github.com/dothiv/translations-updater/csv"
	json "github.com/dothiv/translations-updater/lang/json"
	"github.com/dothiv/translations-updater/util"
	"os"
)

type ImportCommand struct {
	Site *config.Site
}

func NewImportCommand(site *config.Site) (c *ImportCommand) {
	c = new(ImportCommand)
	c.Site = site
	return
}

func (c *ImportCommand) Exec() (err error, errorStrings []csv.KeyError) {
	var csvfile *os.File
	csvfile, err = util.LoadUri(c.Site.Source)
	if err != nil {
		return
	}

	for _, target := range c.Site.Targets {
		r := csv.NewCsvFileReader(csvfile)

		var str map[string]interface{}

		str, err, errorStrings = r.GetStrings(target.Code, target.Val)
		if err != nil {
			return
		}
		var jsfile *os.File
		jsfile, err = os.OpenFile(target.Target, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
		defer jsfile.Close()

		w := json.NewJsonIndentLangWriter()
		w.WriteTo(str, jsfile)

	}

	return
}
