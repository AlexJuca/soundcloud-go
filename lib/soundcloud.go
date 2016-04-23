// Define a type called Client that handles authentication and communication with the sound
// API
package soundclient

import ( "fmt"
          "net/http"
          "io/ioutil"
          "github.com/antonholmquist/jason"
        )

const apiBaseUrl string = "https://api.soundcloud.com"

type SoundCloud struct {
    ClientId string
    ClientSecret string
}

// Handle get requests to an endpoint
// and return a *jason.Object containing data
func Get(s SoundCloud, apiCall string, kargs string) *jason.Object {
    url := apiBaseUrl+"/"+apiCall+"/"+kargs+"?client_id="+s.ClientId
    resp, err := http.Get(url)
    if (err != nil) {
        fmt.Println("[ERROR] ", err)
    } else {
        result, err := ioutil.ReadAll(resp.Body)
        if (err == nil) {
            v, _ := jason.NewObjectFromBytes(result)
            return v
        }
    }
    defer resp.Body.Close()
    return nil
    
}

// Return a soundcloud track 
func (s SoundCloud) Tracks(trackId string) *jason.Object {
    return Get(s, "tracks", trackId)
}

func (s SoundCloud) Playlists(playlistsId string) *jason.Object {
    return Get(s, "playlists", playlistsId)
}


// Return a soundcloud user
func (s SoundCloud) Users(userId string) *jason.Object {
    return Get(s, "users", userId)
}


// Get Group members and contributed tracks
func (s SoundCloud) Groups(groupId string) *jason.Object {
    return Get(s, "groups", groupId)
}
 
// Handle POST requests
func Post() {
    
}




                        

