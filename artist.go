package spotify

//ArtistAPI structure
type ArtistAPI struct {
	ExternalURLs map[string]string `json:"external_urls"`
	Followers    struct {
		Total int `json:"total"`
	} `json:"followers"`
	Genres     []string     `json:"genres"`
	Href       string       `json:"href"`
	ID         string       `json:"id"`
	Images     []AlbumImage `json:"images"`
	Name       string       `json:"name"`
	Popularity int          `json:"popularity"`
	Type       string       `json:"type"`
	URI        string       `json:"uri"`
}
//todo : complete all methods