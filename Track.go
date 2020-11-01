package spotify

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//TrackArtist from get
type TrackArtist struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"` //example : Alec Benjamin
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

//AlbumImage from get
type AlbumImage struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

//AudioFeature struct
type AudioFeature struct {
	Danceability     json.Number `json:"danceability"`
	Energy           json.Number `json:"energy"`
	Key              json.Number `json:"key"`
	Loudness         json.Number `json:"loudness"`
	Mode             json.Number `json:"mode"`
	Speechiness      json.Number `json:"speechiness"`
	Acousticness     json.Number `json:"acousticness"`
	Instrumentalness json.Number `json:"instrumentalness"`
	Liveness         json.Number `json:"liveness"`
	Valence          json.Number `json:"valence"`
	Tempo            json.Number `json:"tempo"`
	Type             string      `json:"type"`
	ID               string      `json:"id"`
	URI              string      `json:"uri"`
	TrackHref        string      `json:"track_href"`
	AnalysisURL      string      `json:"analysis_url"`
	Duration         int         `json:"duration"`
	TimeSignature    int         `json:"time_signature"`
}

//AudioFeatureAPI from get
type AudioFeatureAPI struct {
	AudioFeatures []AudioFeature `json:"audio_features"`
}

//TrackAlbum from get
type TrackAlbum struct {
	AlbumType            string            `json:"album_type"`
	Artists              []TrackArtist     `json:"artists"`
	ExternalUrls         map[string]string `json:"external_urls"`
	Href                 string            `json:"href"`
	ID                   string            `json:"id"`
	Images               []AlbumImage      `json:"images"`
	Name                 string            `json:"name"`
	ReleaseDate          string            `json:"release_date"`
	ReleaseDatePrecision string            `json:"release_date_precision"`
	TotalTracks          int               `json:"total_tracks"`
	Type                 string            `json:"type"`
	URI                  string            `json:"uri"`
}

//TrackAPI2 struct
type TrackAPI2 struct {
	Tracks []TrackAPI `json:"tracks"`
}

//TrackAPI from get
type TrackAPI struct {
	Album        TrackAlbum        `json:"album"`
	Artists      []TrackArtist     `json:"artists"`
	Markets      []string          `json:"available_markets"`
	DiscNumber   int               `json:"disc_number"`
	Duration     int               `json:"duration_ms"`
	Explicit     bool              `json:"explicit"`
	ExternalIDs  map[string]string `json:"external_ids"`
	ExternalUrls map[string]string `json:"external_urls"`
	ID           string            `json:"id"`
	IsLocal      bool              `json:"is_local"`
	Name         string            `json:"name"`
	Popularity   int               `json:"popularity"`
	PreviewURL   string            `json:"preview_url"`
	TrackNumber  int               `json:"track_number"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
	Href         string            `json:"href"`
}

//GetTrack return track from its id
func (cl *Client) GetTrack(id string) (TrackAPI, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/tracks/"+id+"?market=US", nil)
	req.Header.Add("Authorization", "Bearer "+cl.Token)
	res, err := client.Do(req)
	var err2 error
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
	track := TrackAPI{}
	Jsonerr := json.Unmarshal(body, &track)
	if Jsonerr != nil {
		err2 = Jsonerr
	}
	return track, err2
}

//GetTracks returns several tracks
func (cl *Client) GetTracks(id []string) (TrackAPI2, error) {
	client := &http.Client{}
	uri := "https://api.spotify.com/v1/tracks?ids=" + url.QueryEscape(strings.Join(id, ",")) + "&market=US"
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
	tracks := TrackAPI2{}
	Jsonerr := json.Unmarshal(body, &tracks)
	if Jsonerr != nil {
		err2 = Jsonerr
	}
	return tracks, err2
}

//AudioFeatures returns array of tracks provided by thier id
func (cl *Client) AudioFeatures(id []string) (AudioFeatureAPI, error) {
	client := &http.Client{}
	uri := "https://api.spotify.com/v1/audio-features?ids=" + url.QueryEscape(strings.Join(id, ","))
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
	features := AudioFeatureAPI{}
	Jsonerr := json.Unmarshal(body, &features)
	if Jsonerr != nil {
		err2 = Jsonerr
	}
	return features, err2
}

//todo : AudioAnalysis
