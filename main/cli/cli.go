//
// cli bin
//
// Used to manually convert CSV files to JSON lang files
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package main

import (
	"flag"
	"fmt"
	"github.com/dothiv/translations-updater/command"
	"github.com/dothiv/translations-updater/config"
	"os"
)

func main() {
	src := flag.String("source", "", "source CSV file")
	trgt := flag.String("target", "", "target JSON file")
	codeCol := flag.String("code", "Code", "name of the code column")
	valCol := flag.String("val", "Text Deutsch", "name of the value column")
	flag.Parse()

	if len(*trgt) == 0 {
		os.Stderr.WriteString("target is required\n")
		flag.Usage()
		os.Exit(1)
	}

	if len(*src) == 0 {
		os.Stderr.WriteString("source is required\n")
		flag.Usage()
		os.Exit(1)
	}

	os.Stdout.WriteString(fmt.Sprintf("Opening %s ...\n", *src))
	site := new(config.Site)
	site.Source = *src
	target := new(config.Target)
	target.Code = *codeCol
	target.Val = *valCol
	target.Target = *trgt
	site.Targets = append(site.Targets, target)
	c := command.NewImportCommand(site)
	err, errorStrings := c.Exec()
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		for _, v := range errorStrings {
			os.Stderr.WriteString(v.Msg)
		}
		os.Exit(1)
	}
	os.Stdout.WriteString(fmt.Sprintf("%s written.\n", *trgt))
	return
}
