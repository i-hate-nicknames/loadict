package oxford

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nvm.ga/loadict/pkg/load/loader"
)

const urlTemplate = "https://od-api.oxforddictionaries.com/api/v2/entries/en-us/%s" +
	"?strictMatch=false&fields=definitions%%2Cpronunciations%%2Cexamples"

const hitsPerMin = 60

var _ loader.Loader = Loader{}

type Loader struct {
	AppID, AppKey string
}

func (f Loader) Load(word string) (string, error) {
	resp, err := f.fetchWord(word)
	if err != nil {
		return "", err
	}
	result, err := renderCard(resp)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (f Loader) GetRPM() int {
	return hitsPerMin
}

func (f Loader) fetchWord(word string) (*Response, error) {
	client := &http.Client{}
	req, err := assembleRequest(f.AppID, f.AppKey, word)
	if err != nil {
		return nil, fmt.Errorf("cannot intantiate http request: %s", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong response status: %s", resp.Status)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var response Response
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func assembleRequest(appID, appKey, word string) (*http.Request, error) {
	url := fmt.Sprintf(urlTemplate, word)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("app_id", appID)
	req.Header.Add("app_key", appKey)
	return req, nil
}
