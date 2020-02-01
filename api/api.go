package api

import (
	"encoding/json"
	"net/http"
)

// Coord thing
type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Snake thing
type Snake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int     `json:"health"`
	Body   []Coord `json:"body"`
}

// Board thing
type Board struct {
	Height int     `json:"height"`
	Width  int     `json:"width"`
	Food   []Coord `json:"food"`
	Snakes []Snake `json:"snakes"`
}

// Game thing
type Game struct {
	ID string `json:"id"`
}

// SnakeRequest thing
type SnakeRequest struct {
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

// StartResponse thing
type StartResponse struct {
	Color string `json:"color,omitempty"`
}

// MoveResponse thing
type MoveResponse struct {
	Move string `json:"move"`
}

// DecodeSnakeRequest thing
func DecodeSnakeRequest(req *http.Request, decoded *SnakeRequest) error {
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return err
}
