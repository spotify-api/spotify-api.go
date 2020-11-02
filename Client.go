package Spotify

type ClientStruct struct {
	Token   string
	Artists ArtistStruct
	Tracks  TrackStruct
}

func Client(token string) ClientStruct {

	return ClientStruct{
		Token:   token,
		Artists: Artist(token),
		Tracks:  Tracks(token),
	}

}
