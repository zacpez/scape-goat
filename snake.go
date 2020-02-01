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
		return api.RIGHT
	}
	if dx < 0 {
		return api.LEFT
	}
	if dy < 0 {
		return api.DOWN
	}
	return api.UP
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
func ComputeDirection(snake *api.Snake) api.Direction {
	var exclude = []api.Direction{NeckDirection(snake)}
	var options = Difference(api.DirectionChoices, exclude)
	return options[SimpleRandomChoice(len(options))]
}
