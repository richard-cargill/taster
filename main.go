package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shkh/lastfm-go/lastfm"
	"log"
	"net/http"
	"os"
)

type Track struct {
	Name      string `json:"name,omitempty"`
	Artist    string `json:"artist,omitempty"`
	Album     string `json:"album,omitempty"`
	Image     string `json:"image,omitempty"`
	Date      string `json:"date,omitempty"`
	IsPlaying string `json:"isPlaying,omitempty"`
}

var tracks []Track

func GetTracks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(&tracks)
}

func main() {
	port := os.Getenv("PORT")
	apiKey := os.Getenv("APIKEY")
	apiSecret := os.Getenv("APISECRET")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	api := lastfm.New(apiKey, apiSecret)
	result, _ := api.User.GetRecentTracks(lastfm.P{
		"user":  "catdoce",
		"limit": 7,
	})

	router := mux.NewRouter()

	for _, u := range result.Tracks {
		count := len(u.Images)

		track := Track{
			Name:      u.Name,
			Artist:    u.Artist.Name,
			Album:     u.Album.Name,
			Image:     u.Images[count-1].Url,
			Date:      u.Date.Date,
			IsPlaying: u.NowPlaying,
		}
		tracks = append(tracks, track)
	}

	router.HandleFunc("/tracks", GetTracks).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
