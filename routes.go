package main

import (
	"log"
	"net/http"

	"github.com/zacpez/scape-goat/api"
	"github.com/zacpez/scape-goat/snake"
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
		Color: "#69604D",
	})
}

// Move thing
func Move(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}
	var direction = snake.ComputeDirection(&decoded.You, &decoded.Board)

	respond(res, api.MoveResponse{
		Move: string(direction),
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
