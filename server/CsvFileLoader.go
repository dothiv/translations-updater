package server

import (
	"encoding/csv"
	"os"
)

type CsvFileLoader struct {
	Reader *csv.Reader
	Uri    string
}

func NewCsvFileLoader(uri string) (l *CsvFileLoader, err error) {
	l = new(CsvFileLoader)
	l.Uri = uri
	file, err := os.Open(uri)
	if err != nil {
		return
	}
	l.Reader = csv.NewReader(file)
	return
}
