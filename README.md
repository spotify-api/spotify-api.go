# spotify-api.go
A spotify-api wrapper in golang

# Installation
```bash
go get github.com/spotify-api/spotify-api.go
```

# Client
To get the client simply import the module and call the `New` function
```go
import (
 "github.com/spotify-api/spotify-api.go"
)
client := spotify.New("token") //the oauth token
```

# Search
Searching artists or tracks
```go
import (
 "fmt"
 "github.com/spotify-api/spotify-api.go"
)
client := spotify.New("token") //the oauth token
tracks,err := client.SearchTracks("oh my god by alec benjamin")
if err != nil {
 fmt.Println("ERROR: " + err) //incase if there is an Error
 return
} 
fmt.Println(tracks.Tracks.Items[0].Name) // Prints : Oh My God

//Searching for artists
artists ,err2 := client.SearchArtists("Alec Benjamin")
if err2 != nil {
 fmt.Println("ERROR: " + err) //incase if there is an Error
 return
} 
fmt.Println(artists.Artists.Items[0].Name) // Prints : Alec Benjamin
```
