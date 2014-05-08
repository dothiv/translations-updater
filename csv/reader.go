//
// CsvFileReader
//
// Use to read CSV files and convert them to our JSON format.
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// Allowed format for keys.
var KEY_FORMAT = regexp.MustCompile(`^([a-zA-Z][a-zA-Z0-9]*\.)*[a-zA-Z][a-zA-Z0-9]*$`)

// Error type used when data in the CSV contains errors
var CsvDataError = errors.New("CSV data error")

// used to describe CSV data errors
type KeyError struct {
	Key string
	Msg string
}

// Creates a new CSV data error
func NewKeyError(key string, msg string) (e *KeyError) {
	e = new(KeyError)
	e.Key = key
	e.Msg = msg
	return
}

// This is the main object
type CsvFileReader struct {
	file io.Reader
}

// Creates a new CsvFileReader for a io.Reader
func NewCsvFileReader(file io.Reader) (r *CsvFileReader) {
	r = new(CsvFileReader)
	r.file = file
	return
}

// Returns a map of strins from the CSV file using keyCol value as the name of the key and valCol as the column for the value
// The method builds a deep map of strings:
// "this.is.a.key,value" turns into
// str["this"]["is"]["a"]["key"] = "value"
func (r *CsvFileReader) GetStrings(keyCol string, valCol string) (str map[string]interface{}, err error, dataErrors []KeyError) {
	reader := csv.NewReader(r.file)
	keyColIndex, valColIndex, colError := r.getColIndex(reader, keyCol, valCol)
	if colError != nil {
		err = colError
		return
	}
	lines, readErr := reader.ReadAll()
	if readErr != nil {
		err = readErr
		return
	}
	str = make(map[string]interface{})
	for _, v := range lines {
		key := v[keyColIndex]
		if !KEY_FORMAT.MatchString(key) {
			e := NewKeyError(key, "The key \""+key+"\" is invalid. Please only use A-z, 0-9, dots for seperation and no numbers in the beginning or after dots.")
			dataErrors = append(dataErrors, *e)
		}
		path := strings.Split(key, ".")
		npaths := len(path)
		parent := str
		for i, part := range path {
			hasChildren := i < npaths-1
			children, exists := parent[part]
			_, childrenIsString := children.(string)
			if exists {
				if hasChildren {
					if childrenIsString {
						// Error: Overlap
						e := NewKeyError(key, "A prefix of the key \""+key+"\" is already in use.")
						dataErrors = append(dataErrors, *e)
					} else {
						parent = parent[part].(map[string]interface{})
					}
				} else {
					// Error: Key in use
					e := NewKeyError(key, "The key \""+key+"\" is used twice.")
					dataErrors = append(dataErrors, *e)
				}
			} else {
				if hasChildren {
					parent[part] = make(map[string]interface{})
					parent = parent[part].(map[string]interface{})
				} else {
					parent[part] = v[valColIndex]
				}
			}
		}
	}
	if len(dataErrors) > 0 {
		err = CsvDataError
	}
	return
}

// Searches the first line for the given names and returns the column indices
func (r *CsvFileReader) getColIndex(reader *csv.Reader, keyColName string, valColName string) (keyColIndex int, valColIndex int, err error) {
	line, err := reader.Read()
	if err != nil {
		return
	}
	for k, v := range line {
		if v == keyColName {
			keyColIndex = k
		}
		if v == valColName {
			valColIndex = k
		}
	}
	if keyColIndex >= 0 && valColIndex >= 0 {
		return
	}
	err = fmt.Errorf("Failed to find index for columns '%s' and '%s!", keyColName, valColName)
	return
}
