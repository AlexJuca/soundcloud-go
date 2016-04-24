package soundclient

import (
        "testing" 
        "os"
        )

// NOTE THAT tests will fail if you do not insert the proper CLIENT_ID when initialising the SoundCloud object
// Add the clientId to your enviroment 
func makeObj() SoundCloud {
	return SoundCloud{os.Getenv("SOUNDCLOUD_CLIENT_ID"), os.Getenv("SOUNDCLOUD_SECRET")}
}

func TestGet(t *testing.T) {
	s := makeObj()
	Get(s, "tracks", "12121")
}

func TestTracks(t *testing.T) {
	s := makeObj()
	tTrack := s.Tracks("13158")
	title, _ := tTrack.GetString("title")
	if title == "One (Playmaker Remix)" {

	} else {
		t.Errorf("Did not get the correct song title")
	}
}

func TestGroup(t *testing.T) {
	s := makeObj()
	tGroup := s.Groups("3")
	uri, _ := tGroup.GetString("uri")
	if uri != "https://api.soundcloud.com/groups/3" {
		t.Errorf("Unable to get correct uri for group")
	}

}
