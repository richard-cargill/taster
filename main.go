package main

import (
	"encoding/json"
	"fmt"
	"github.com/shkh/lastfm-go/lastfm"
)

type Track struct {
	Name   string `json:"name,omitempty"`
	Artist string `json:"artist,omitempty"`
	Album  string `json:"album,omitempty"`
	Image  string `json:"image,omitempty"`
}

type Tracks []Track

var tracks Tracks

func main() {
	api := lastfm.New("", "")

	result, _ := api.User.GetRecentTracks(lastfm.P{"user": "catdoce"})

	for _, u := range result.Tracks {
		count := len(u.Images)

		track := Track{
			Name:   u.Name,
			Artist: u.Artist.Name,
			Album:  u.Album.Name,
			Image:  u.Images[count-1].Url,
		}
		tracks = append(tracks, track)
	}

	response, _ := json.Marshal(&tracks)
	fmt.Println(string(response))
}
