package Spotify

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type FetchOptions struct {
	token  string
	link   string
	params string
}

func fetch(options FetchOptions, structure interface{}) error {
	reqClient := &http.Client{}

	req, err := http.NewRequest("GET", (baseUrl + options.link + "?" + options.params), nil)
	req.Header.Add("Authorization", ("Bearer " + options.token))

	res, err := reqClient.Do(req)
	body, readErr := ioutil.ReadAll(res.Body)
	var fetchErr error

	if readErr != nil {
		fetchErr = readErr
	}
	if err != nil {
		fetchErr = errors.New("Unexpected Error: Request failed")
	}
	if res.StatusCode != 200 {
		fetchErr = errors.New("Unexpected Error: Probably failiure of api!")
	}

	jsonErr := json.Unmarshal(body, structure)

	if jsonErr != nil {
		fetchErr = jsonErr
	}
	return fetchErr
}
