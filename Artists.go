package Spotify

type ArtistStruct struct {
	Get func(id string) (ArtistSchema, error)
}

func Artist(token string) ArtistStruct {

	return ArtistStruct{

		Get: func(id string) (ArtistSchema, error) {
			var result ArtistSchema

			err := fetch(FetchOptions{
				token:  token,
				link:   "artists/" + id,
				params: "",
			}, &result)

			return result, err
		},
	}

}
