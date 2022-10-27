package main

import (
	"encoding/json"
	"net/http"
)

type AppData struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

type app struct{}

func (*app) GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := &AppData{
		SlackUsername: "Oreoluwa10",
		Backend:       false,
		Age:           34,
		Bio:           "talk to me nice....",
	}
	json.NewEncoder(w).Encode(data)
}

func newApp() *app {
	return &app{}
}
