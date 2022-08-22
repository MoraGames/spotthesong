package main

import (
	spotifyauth "github.com/zmb3/spotify/auth"
	"log"
	"os"
)

const (
	RedirectURL = "spotthesong.altervista.org"
	SessionID = "spotthesong"
)

var (
	SpotifyAuth *spotifyauth.Authenticator = nil
	SpotifyUrl string = ""
)

func init(){
	var c Config
	c.GetConfig()

	os.Setenv("SPOTIFY_ID", c.SPOTIFY_ID)
	os.Setenv("SPOTIFY_SECRET", c.SPOTIFY_SECRET)

	SpotifyAuth = spotifyauth.New(spotifyauth.WithRedirectURL(RedirectURL), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))
	SpotifyUrl = SpotifyAuth.AuthURL(SessionID)
}

func main(){
	log.Printf("[ Running ]\n")
	//mainRouter := mux.NewRouter()
	//mainRouter.HandleFunc("/", handler.HomePageHandler)
}