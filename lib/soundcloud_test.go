package soundclient

import "testing"


// NOTE THAT tests will fail if you do not insert the proper CLIENT_ID when initialising the SoundCloud object
func makeObj() SoundCloud {
    return SoundCloud{"042bd096125d480ea6d7f08cf4b092eb", "042bd096125d480ea6d7f08cf4b092eb"}
}

func TestGet(t *testing.T) {
    s := makeObj()
    Get(s, "tracks", "12121")
}


func TestTracks(t *testing.T) {
    s := makeObj()
    tTrack := s.Tracks("13158")
    title, _ := tTrack.GetString("title")
    if (title == "One (Playmaker Remix)") {
        
    } else {
        t.Errorf("Did not get the correct song title")
    }
}

func TestGroup(t *testing.T) {
    s := makeObj()
    tGroup := s.Groups("3")
    uri, _ := tGroup.GetString("uri")
    if (uri != "https://api.soundcloud.com/groups/3"){
        t.Errorf("Unable to get correct uri for group")
    }
    
}