package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const urlTemplate = "https://od-api.oxforddictionaries.com/api/v2/entries/en-us/%s?strictMatch=false&fields=definitions%%2Cexamples"

func fetchWord(appId, appKey, word string) (*Response, error) {
	var response Response
	client := &http.Client{}
	url := fmt.Sprintf(urlTemplate, word)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("app_id", appId)
	req.Header.Add("app_key", appKey)
	log.Println("performing request to " + url)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong response status: %s", resp.Status)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
