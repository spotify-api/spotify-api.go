package Spotify

import "strings"

type TrackStruct struct {
	Get           func(id string) (TrackAPI, error)
	AudioFeatures func(ids []string) (AudioFeatureAPI, error)
}

func Tracks(token string) TrackStruct {

	return TrackStruct{

		Get: func(id string) (TrackAPI, error) {
			var result TrackAPI

			err := fetch(FetchOptions{
				token:  token,
				link:   "tracks/" + id,
				params: "",
			}, &result)

			return result, err
		},

		AudioFeatures: func(ids []string) (AudioFeatureAPI, error) {
			var result AudioFeatureAPI

			err := fetch(FetchOptions{
				token:  token,
				link:   ("tracks/" + strings.Join(ids, ",")),
				params: "",
			}, &result)

			return result, err
		},
	}

}
