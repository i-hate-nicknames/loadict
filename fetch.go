package main

import (
	"fmt"
	"net/http"
	"time"
)

const urlTemplate = "https://od-api.oxforddictionaries.com/api/v2/entries/en-us/%s?strictMatch=false&fields=definitions%%2Cpronunciations%%2Cexamples"

func fetchWord(appId, appKey, word string) (*Response, error) {
	time.Sleep(2 * time.Second)
	// client := &http.Client{}
	// req, err := assembleRequest(appId, appKey, word)
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot intantiate http request: %s", err)
	// }
	// log.Println("performing request to " + req.RequestURI)
	// resp, err := client.Do(req)
	// if err != nil {
	// 	return nil, err
	// }
	// if resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("Wrong response status: %s", resp.Status)
	// }
	// defer resp.Body.Close()
	// decoder := json.NewDecoder(resp.Body)
	var response Response
	// err = decoder.Decode(&response)
	// if err != nil {
	// 	return nil, err
	// }
	return &response, nil
}

func assembleRequest(appId, appKey, word string) (*http.Request, error) {
	url := fmt.Sprintf(urlTemplate, word)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("app_id", appId)
	req.Header.Add("app_key", appKey)
	return req, nil
}
