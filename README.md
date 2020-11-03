<div align="center">
  <img src="https://media.discordapp.net/attachments/736466510888960020/760853915876327464/Sa.png?width=718&height=275"><br>
  <div>
    <a href="https://spotify-apijs.netlify.app/#/"><img src="https://img.shields.io/badge/READ-DOCS-orange?style=for-the-badge"></a>
    <a href="https://github.com/spotify-api/spotify-api.js/"><img src="https://img.shields.io/github/repo-size/spotify-api/spotify-api.js?label=Size&style=for-the-badge"></a>
    <a href="https://www.npmjs.com/package/spotify-api.js"><img src="https://img.shields.io/npm/v/spotify-api.js?label=Version&style=for-the-badge"></a>
  </div><br>
</div>

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
client := spotify.Client("token") //the oauth token
```

# Search
Searching artists or tracks
```go
import (
 "fmt"
 "github.com/spotify-api/spotify-api.go"
)
client := spotify.Client("token") //the oauth token
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
