package main

import (
	"fmt"
	"github.com/dothiv/translations-updater/csv"
	json "github.com/dothiv/translations-updater/lang/json"
	"os"
)

func main() {
	src := os.Args[1]
	target := os.Args[2]
	codeCol := os.Args[3]
	valCol := os.Args[4]

	os.Stdout.WriteString(fmt.Sprintf("Source: %s\n", src))
	os.Stdout.WriteString(fmt.Sprintf("Target: %s\n", target))
	os.Stdout.WriteString(fmt.Sprintf("Codes:  %s\n", codeCol))
	os.Stdout.WriteString(fmt.Sprintf("Values: %s\n", valCol))

	csvfile, _ := os.Open(src)
	defer csvfile.Close()
	r := csv.NewCsvFileReader(csvfile)
	str, err, errorStrings := r.GetStrings(codeCol, valCol)
	if err != nil {
		for _, v := range errorStrings {
			os.Stdout.WriteString(v.Msg)
		}
		os.Exit(1)
	}

	jsfile, _ := os.OpenFile(target, os.O_RDWR|os.O_TRUNC, 0644)
	defer jsfile.Close()
	w := json.NewJsonLangWriter()
	w.WriteTo(str, jsfile)
}
