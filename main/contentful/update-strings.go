//
// update-strings bin
//
// Used to update the strings on contentful with the strings from the CSV file
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dothiv/contentful/model"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	src := flag.String("source", "", "source CSV file")
	space := flag.String("space", "", "space id")
	token := flag.String("token", "", "delivery api token")
	contentToken := flag.String("ctoken", "", "content management api token")
	ctype := flag.String("type", "", "ID of content type")
	codeCol := flag.String("code", "Code", "name of the code column")
	valEnCol := flag.String("en", "Text Englisch", "name of the value column for the english string")
	valDeCol := flag.String("de", "Text Deutsch", "name of the value column for german string")

	flag.Parse()

	csvFile, _ := os.Open(*src)
	reader := csv.NewReader(csvFile)
	reader.TrailingComma = true

	line, err := reader.Read()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	// Get colIndex
	keyColIndex, keyColError := GetColIndex(line, *codeCol)
	if keyColError != nil {
		os.Stderr.WriteString(keyColError.Error())
		os.Exit(1)
	}
	valEnColIndex, valEnColError := GetColIndex(line, *valEnCol)
	if valEnColError != nil {
		os.Stderr.WriteString(valEnColError.Error())
		os.Exit(1)
	}
	valDeColIndex, valDeColError := GetColIndex(line, *valDeCol)
	if valDeColError != nil {
		os.Stderr.WriteString(valDeColError.Error())
		os.Exit(1)
	}

	// Read lines
	lines, readErr := reader.ReadAll()
	if readErr != nil {
		os.Stderr.WriteString(readErr.Error())
		os.Exit(1)
	}

	client := &http.Client{}

	for _, textLine := range lines {
		code := strings.TrimSpace(textLine[keyColIndex])
		if len(code) == 0 {
			continue
		}
		en := textLine[valEnColIndex]
		de := textLine[valDeColIndex]
		os.Stdout.WriteString(fmt.Sprintf("%s:\n", code))
		os.Stdout.WriteString(fmt.Sprintf("en: %s\n", en))
		os.Stdout.WriteString(fmt.Sprintf("de: %s\n", de))
		// Find entry
		uri := fmt.Sprintf("https://cdn.contentful.com/spaces/%s/entries?content_type=%s&fields.code=%s", *space, *ctype, code)
		req, _ := http.NewRequest("GET", uri, nil)
		req.Header.Add("Content-Type", "application/vnd.contentful.management.v1+json")
		req.Header.Add("Authorization", "Bearer "+*token)
		resp, respErr := client.Do(req)
		if respErr != nil {
			os.Stderr.WriteString(respErr.Error())
			os.Exit(1)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		var data model.SearchResponse
		jsonErr := json.Unmarshal(body, &data)
		if jsonErr != nil {
			os.Stderr.WriteString(jsonErr.Error())
			os.Exit(1)
		}
		if data.Sys.Type == "Error" {
			os.Stderr.WriteString(data.Message)
			os.Exit(1)
		}

		if data.Total == 0 {
			// New item
		} else if data.Total == 1 {
			// Update item
			os.Stdout.WriteString(fmt.Sprintf("Updating: %s\n", data.Items[0].Fields["code"]))
			item := fetchItem(*space, *contentToken, data.Items[0].Sys.Id)
			item.Fields.Value.En = en
			item.Fields.Value.De = de
			updatedEntry, _ := json.Marshal(item.Fields)
			os.Stdout.Write(updatedEntry)
			os.Exit(2)
		} else {
			// Multiple items found
			os.Stderr.WriteString(fmt.Sprintf("Multiple entries for '%s' found!", code))
		}
	}
}

func fetchItem(spaceId string, token string, id string) (item model.Item) {

	uri := fmt.Sprintf("https://api.contentful.com/spaces/%s/entries/%s", spaceId, id)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Content-Type", "application/vnd.contentful.management.v1+json")
	req.Header.Add("Authorization", "Bearer "+token)
	resp, respErr := client.Do(req)
	if respErr != nil {
		os.Stderr.WriteString(respErr.Error())
		os.Exit(1)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	jsonErr := json.Unmarshal(body, &item)
	if jsonErr != nil {
		os.Stderr.WriteString(jsonErr.Error())
		os.Exit(1)
	}
	return
}

// Searches the first line for the given names and returns the column indices
func GetColIndex(line []string, colName string) (colIndex int, err error) {
	for k, v := range line {
		if v == colName {
			colIndex = k
		}
	}
	if colIndex >= 0 {
		return
	}
	err = fmt.Errorf("Failed to find index for column '%s'!", colIndex)
	return
}
