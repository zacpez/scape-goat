package snake

import (
	"sync"
	"time"

	"github.com/zacpez/scape-goat/api"
)

// DumbDirections finds a []Direction to exclude from choices
func DumbDirections(respone chan<- []api.Direction, snake *api.Snake, board *api.Board, directions *[]api.Direction) []api.Direction {
	var dumbIdeas = []api.Direction{}
	var channelCount = len(*directions)
	var wg sync.WaitGroup

	// Setup Multiple thread channels
	wg.Add(channelCount)
	badies := make(chan api.Direction, channelCount)

	for _, direction := range *directions {
		go PredictNextDirection(badies, &wg, snake, board, &direction)
	}

	wg.Wait()
	close(badies)

	for badie := range badies {
		dumbIdeas = append(dumbIdeas, badie)
	}

	return dumbIdeas
}

// PredictNextDirection returns bad decissions
func PredictNextDirection(badies chan<- api.Direction, wg *sync.WaitGroup, snake *api.Snake, board *api.Board, direction *api.Direction) {
	var dx = snake.Body[0].X
	var dy = snake.Body[0].Y
	if *direction == api.DOWN {
		dy = snake.Body[0].Y + 1
	}
	if *direction == api.UP {
		dy = snake.Body[0].Y - 1
	}
	if *direction == api.LEFT {
		dy = snake.Body[0].Y - 1
	}
	if *direction == api.RIGHT {
		dy = snake.Body[0].Y + 1
	}

	for _, other := range board.Snakes {
		for _, part := range other.Body {
			if part.X == dx || part.Y == dy {
				badies <- *direction
			}
		}
	}

	wg.Done()
}

// ComputeDirection thing
func ComputeDirection(snake *api.Snake, board *api.Board) api.Direction {
	// Quickly find the worst options
	var excludeChoices = []api.Direction{
		NeckDirection(snake),
		EdgeDirection(snake, board.Width, board.Height),
		api.NONE}
	var choices = Difference(api.DirectionChoices, excludeChoices)

	// Find more logical directions
	respond := make(chan []api.Direction, 1)
	go DumbDirections(respond, snake, board, &choices)

	select {
	case dumbIdeas := <-respond:
		// Exclude the likely bad options from DirectionChoices
		choices = Difference(api.DirectionChoices, dumbIdeas)
	case <-time.After(api.TIMEOUT):
	}

	var choiceCount = len(choices)
	// One choice left
	if choiceCount == 1 {
		return choices[0]
	}
	// Choose a random choice
	return choices[SimpleRandomChoice(choiceCount)]
}
