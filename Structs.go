package spotify

import (
	"encoding/json"
)

var baseUrl string = "https://api.spotify.com/v1/"

type Followers struct {
	Total int `json:"total"`
}

type ArtistSchema struct {
	ExternalURLs map[string]string `json:"external_urls"`
	Followers    Followers         `json:"followers"`
	Genres       []string          `json:"genres"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Popularity   int               `json:"popularity"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type TrackArtist struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type AlbumImage struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

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

type AudioFeatureAPI struct {
	AudioFeatures []AudioFeature `json:"audio_features"`
}

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

type TrackAPI2 struct {
	Tracks []TrackAPI `json:"tracks"`
}

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
