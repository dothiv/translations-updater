package google

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func WatchDocument(ident string, token string, webhook_address string) (err error) {
	uri := fmt.Sprintf("https://www.googleapis.com/drive/v2/files/%s/watch", ident)
	body := make(map[string]string)
	// See https://developers.google.com/drive/web/push for a description of the fields
	body["id"] = makeUUID()
	body["type"] = "web_hook"
	body["address"] = webhook_address

	b, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", uri, bytes.NewReader(b))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respons_body, err := ioutil.ReadAll(resp.Body)
	os.Stdout.Write(respons_body)
	return
}

func makeUUID() string {
	return "some-uuid"
}
