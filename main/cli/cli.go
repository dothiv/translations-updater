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
	"github.com/dothiv/translations-updater/csv"
	json "github.com/dothiv/translations-updater/lang/json"
	"os"
)

func main() {
	src := flag.String("source", "", "source CSV file")
	target := flag.String("target", "", "target JSON file")
	codeCol := flag.String("code", "Code", "name of the code column")
	valCol := flag.String("val", "Text Deutsch", "name of the value column")
	flag.Parse()

	if len(*target) == 0 {
		os.Stderr.WriteString("target is required\n")
		flag.Usage()
		os.Exit(1)
	}

	if len(*src) == 0 {
		os.Stderr.WriteString("source is required\n")
		flag.Usage()
		os.Exit(1)
	}

	csvfile := open(*src)
	r := csv.NewCsvFileReader(csvfile)
	str, err, errorStrings := r.GetStrings(*codeCol, *valCol)
	if err != nil {
		for _, v := range errorStrings {
			os.Stdout.WriteString(v.Msg)
		}
		os.Exit(1)
	}

	jsfile := openFile(*target, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	w := json.NewJsonIndentLangWriter()
	w.WriteTo(str, jsfile)
	return
}

func open(filename string) (f *os.File) {
	return openFile(filename, os.O_RDONLY, 0644)
}

func openFile(filename string, flags int, mode os.FileMode) (f *os.File) {
	f, err := os.OpenFile(filename, flags, mode)
	if err == nil {
		return
	}
	if !os.IsExist(err) {
		os.Stderr.WriteString(fmt.Sprintf("Failed to open file '%s'!\n", filename))
	} else {
		os.Stderr.WriteString(err.Error() + "\n")
	}
	os.Exit(1)

	defer f.Close()
	return
}
