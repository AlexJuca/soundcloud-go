// Define a type called Client that handles authentication and communication with the sound
// API
package soundclient

import ( "fmt"
          "net/http"
          "io/ioutil"
        )

const apiBaseUrl string = "https://api.soundcloud.com"

type SoundCloud struct {
    ClientId string
    ClientSecret string
}

// Handle get requests to an endpoint
// 
func Get(s SoundCloud, apiCall string, kargs string) {
    url := apiBaseUrl+"/"+apiCall+"/"+kargs+"?client_id="+s.ClientId
    resp, err := http.Get(url)
    fmt.Println("[DEBUG] URL-> ", url)
    if (err != nil) {
        fmt.Println("[ERROR] ", err)
    } else {
        body, err := ioutil.ReadAll(resp.Body)
        if (err == nil) {
            fmt.Printf("%s" ,body)
        }
    }
    defer resp.Body.Close()
    
}


// Return a soundcloud track 
func (s SoundCloud) Tracks(trackId string) {
    Get(s, "tracks", trackId)
}

func (s SoundCloud) Playlists(playlistsId string) {
    Get(s, "playlists", playlistsId)
}



// Handle POST requests
func Post() {
    
}




                        

