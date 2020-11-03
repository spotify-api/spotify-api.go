package spotify

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type AuthRes struct {
	Token string `json:"access_token"`
}

type AuthStruct struct {
	Get func(ID string, secret string) (string, error)
}

func Auth() AuthStruct {

	return AuthStruct{

		Get: func(ID string, secret string) (string, error) {
			var err2 error
			res, Fetcherr := http.PostForm("https://accounts.spotify.com/api/token", url.Values{"grant_type": {"client_credentials"}, "client_id": {ID}, "client_secret": {secret}})

			if Fetcherr != nil {
				err2 = Fetcherr
			}
			if res.StatusCode != 200 {
				err2 = errors.New("Spotify API Error (" + strconv.Itoa(res.StatusCode) + ") : " + CleanResponseCode(res.StatusCode))
			}

			body, readErr := ioutil.ReadAll(res.Body)
			if readErr != nil {
				err2 = readErr
			}

			structure := AuthRes{}
			jsonErr := json.Unmarshal(body, &structure)
			if jsonErr != nil {
				err2 = jsonErr
			}
			return structure.Token, err2
		},
	}

}
