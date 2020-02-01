package main

import (
	"time"

	"github.com/zacpez/scape-goat/api"
)

// Head thing
func Head() bool {
	return true
}

// NeckDirection thing
func NeckDirection(snake *api.Snake) api.Direction {
	var dx = snake.Body[0].X - snake.Body[1].X
	var dy = snake.Body[0].Y - snake.Body[1].Y
	if dx > 0 {
		return api.LEFT
	}
	if dx < 0 {
		return api.RIGHT
	}
	if dy < 0 {
		return api.DOWN
	}
	return api.UP
}

// EdgeDirection thing
func EdgeDirection(snake *api.Snake, boardSize int) api.Direction {
	if snake.Body[0].X == 0 {
		return api.LEFT
	}
	if snake.Body[0].X == boardSize {
		return api.RIGHT
	}
	if snake.Body[0].Y == 0 {
		return api.UP
	}
	if snake.Body[0].Y == boardSize {
		return api.DOWN
	}
	return api.NONE
}

// DumbDirections thing
func DumbDirections(snake *api.Snake, board *api.Board) []api.Direction {
	var dumbIdeas = []api.Direction{}
	return dumbIdeas
}

// Difference thing: A - B
func Difference(a []api.Direction, b []api.Direction) (diff []api.Direction) {
	m := make(map[api.Direction]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

// SimpleRandomChoice thing
func SimpleRandomChoice(len int) int {
	return (int(time.Now().UnixNano()) % len)
}

//ComputeDirection thing
func ComputeDirection(snake *api.Snake, board *api.Board) api.Direction {
	var exclude = []api.Direction{
		NeckDirection(snake),
		EdgeDirection(snake, 10),
		api.NONE}
	var dumbIdeas = DumbDirections(snake, board)
	var bestChoices = append(exclude, dumbIdeas...)

	var options = Difference(api.DirectionChoices, bestChoices)
	return options[SimpleRandomChoice(len(options))]
}
