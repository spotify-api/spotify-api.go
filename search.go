package spotify

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	//"strings"
)

//TrackAPISearch structure
type TrackAPISearch struct {
	Tracks struct {
		Items []TrackAPI `json:"items"`
	} `json:"tracks"`
}

//ArtistAPISearch structure
type ArtistAPISearch struct {
	Artists struct {
		Items []ArtistAPI `json:"items"`
	} `json:"artists"`
}

//SearchTracks searches tracks
func (cl *Client) SearchTracks(query string) (TrackAPISearch, error) {
	uri := "https://api.spotify.com/v1/search?q=" + url.QueryEscape(query) + "&market=US&type=track"
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", "Bearer "+cl.Token)
	var err2 error
	res, err := client.Do(req)
	if res.StatusCode != 200 {
		err2 = errors.New("Unexpected Error : Response code is Not 200 OK")
	}
	if err != nil {
		err2 = err
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		err2 = readErr
	}
	tracks := TrackAPISearch{}
	Jsonerr := json.Unmarshal(body, &tracks)
	if Jsonerr != nil {
		err2 = Jsonerr
	}
	return tracks, err2
}

//SearchArtists searches artists
func (cl *Client) SearchArtists(query string) (ArtistAPISearch, error) {
	uri := "https://api.spotify.com/v1/search?q=" + url.QueryEscape(query) + "&market=US&type=artist"
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", "Bearer "+cl.Token)
	var err2 error
	res, err := client.Do(req)
	if res.StatusCode != 200 {
		err2 = errors.New("Unexpected Error : Response code is Not 200 OK")
	}
	if err != nil {
		err2 = err
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		err2 = readErr
	}
	artists := ArtistAPISearch{}
	Jsonerr := json.Unmarshal(body, &artists)
	if Jsonerr != nil {
		err2 = Jsonerr
	}
	return artists, err2
}
