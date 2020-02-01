package main

import (
	"log"
	"net/http"

	"github.com/battlesnakeio/starter-snake-go/api"
)

// Index thing
func Index(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Battlesnake documentation can be found at <a href=\"https://docs.battlesnake.io\">https://docs.battlesnake.io</a>."))
}

// Start thing
func Start(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad start request: %v", err)
	}
	dump(decoded)

	respond(res, api.StartResponse{
		Color: "#75CEDD",
	})
}

// Move thing
func Move(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}
	dump(decoded)

	respond(res, api.MoveResponse{
		Move: "down",
	})
}

// End thing
func End(res http.ResponseWriter, req *http.Request) {
	return
}

// Ping thing
func Ping(res http.ResponseWriter, req *http.Request) {
	return
}
