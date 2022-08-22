package main

import (
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"net/http"
)

var (
	SpotifyToken *oauth2.Token = nil
	SpotifyClient *spotify.Client = nil
)

// the user will eventually be redirected back to your redirect URL
// typically you'll have a handler set up like the following:
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	//Generate the token
	SpotifyToken, err := SpotifyAuth.Token(r.Context(), SessionID, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusNotFound)
		return
	}

	//Create a client using the specified token
	SpotifyClient = spotify.New(SpotifyAuth.Client(r.Context(), SpotifyToken))
}