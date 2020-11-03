package spotify

type ClientStruct struct {
	Token   string
	Artists ArtistStruct
	Tracks  TrackStruct
	Auth    AuthStruct
}

func Client(token string) ClientStruct {

	return ClientStruct{
		Token:   token,
		Artists: Artist(token),
		Tracks:  Tracks(token),
		Auth:    Auth(),
	}

}
