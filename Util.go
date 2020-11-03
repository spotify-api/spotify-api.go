package spotify

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
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
		fetchErr = errors.New("Spotify API Error (" + strconv.Itoa(res.StatusCode) + ") : " + CleanResponseCode(res.StatusCode))
	}

	jsonErr := json.Unmarshal(body, structure)

	if jsonErr != nil {
		fetchErr = jsonErr
	}
	return fetchErr
}

func CleanResponseCode(code int) string {
	knownCodes := []int{201, 202, 204, 304, 400, 401, 403, 404, 429, 500, 502, 503}
	responses := []string{"Created", "Accepted", "No Content", "Not Modified", "Bad Request", "Unauthorized", "Forbidden", "Not Found", "Too Many Requests", "Internal Server Error", "Bad Gateway", "Service Unavailable"}
	index := indexOf(code, knownCodes)
	return responses[index]
}

func indexOf(element int, data []int) int {
	for i, l := range data {
		if element == l {
			return i
		}
	}
	return -1
}
