package main

import (
	"fmt"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	BOARD_WIDTH  = 44
	BOARD_HEIGHT = 20
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

const (
	BLANK_CELL = iota
	SNAKE_CELL
	PELLET_CELL
)

type TickMsg time.Time

type coord struct {
	x int
	y int
}

type snakeBody struct {
	// first element in list is the head of the snake
	body      []coord
	direction int
}

// this is the model used by bubbletea
type snakeGame struct {
	gameBoard [][]int
	snake     snakeBody
	pellet    coord
	score     int
	gameOver  bool
}

func newSnakeGame() snakeGame {
	rand.Seed(time.Now().UnixNano())

	// create a snake in the middle of the board
	snake := snakeBody{
		body: []coord{
			{x: (BOARD_WIDTH / 2) + 1, y: BOARD_HEIGHT / 2},
			{x: BOARD_WIDTH / 2, y: BOARD_HEIGHT / 2},
			{x: (BOARD_WIDTH / 2) - 1, y: BOARD_HEIGHT / 2},
		},
		direction: RIGHT,
	}

	game := snakeGame{
		snake:    snake,
		gameOver: false,
	}

	game.spawnPellet()
	game.updateBoard()

	return game
}

func (s snakeBody) coordInBody(c coord) bool {
	for _, snakePart := range s.body {
		if snakePart == c {
			return true
		}
	}

	return false
}

func (s snakeBody) isHeadCollidingWithBody() bool {
	head := s.body[0]

	for i := 1; i < len(s.body); i++ {
		if head == s.body[i] {
			return true
		}
	}

	return false
}

func (s *snakeGame) updateBoard() {
	gameBoard := [][]int{}

	for i := 0; i < BOARD_HEIGHT; i++ {
		row := []int{}
		for j := 0; j < BOARD_WIDTH; j++ {
			cellType := BLANK_CELL
			curCell := coord{j, i}

			if s.snake.coordInBody(curCell) {
				cellType = SNAKE_CELL
			} else if curCell == s.pellet {
				cellType = PELLET_CELL
			}

			row = append(row, cellType)
		}

		gameBoard = append(gameBoard, row)
	}

	s.gameBoard = gameBoard
}

// place pellet at random position that doesn't overlap with snake
func (s *snakeGame) spawnPellet() {
	pellet := coord{
		x: rand.Intn(BOARD_WIDTH),
		y: rand.Intn(BOARD_HEIGHT),
	}

	for s.snake.coordInBody(pellet) {
		pellet = coord{
			x: rand.Intn(BOARD_WIDTH),
			y: rand.Intn(BOARD_HEIGHT),
		}
	}

	s.pellet = pellet
}

func tickEvery() tea.Cmd {
	return tea.Every(time.Millisecond*100, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (s snakeGame) Init() tea.Cmd {
	return tickEvery()
}

func (s snakeGame) View() string {
	scoreLabel := scoreStyle.Render("score")
	scoreText := fmt.Sprintf("\n%s: %d\n\n", scoreLabel, s.score)

	if s.gameOver {
		return gameOverStyle.Render(gameOverText) + scoreText + "esc to quit\n"
	}

	screen := ""

	for i := 0; i < BOARD_HEIGHT; i++ {
		for j := 0; j < BOARD_WIDTH; j++ {
			if s.gameBoard[i][j] == SNAKE_CELL {
				screen += snakeStyle.Render(" ")
			} else if s.gameBoard[i][j] == PELLET_CELL {
				screen += pelletStyle.Render(" ")
			} else {
				screen += baseStyle.Render(" ")
			}
		}

		if i != BOARD_HEIGHT-1 {
			screen += "\n"
		}
	}

	helpMsg := "arrow keys to move\nesc to quit\n"

	return boardStyle.Render(screen) + scoreText + helpMsg
}

func (s snakeGame) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "esc":
			return s, tea.Quit

		// when moving snake, don't allow movement in opposite direction
		// this is done by checking if snake body part directly behind head
		// is in the direction player is trying to move
		case "up":
			upCoord := coord{x: s.snake.body[0].x, y: s.snake.body[0].y - 1}
			if upCoord != s.snake.body[1] {
				s.snake.direction = UP
			}
		case "right":
			rightCoord := coord{x: s.snake.body[0].x + 1, y: s.snake.body[0].y}
			if rightCoord != s.snake.body[1] {
				s.snake.direction = RIGHT
			}
		case "down":
			downCoord := coord{x: s.snake.body[0].x, y: s.snake.body[0].y + 1}
			if downCoord != s.snake.body[1] {
				s.snake.direction = DOWN
			}
		case "left":
			leftCoord := coord{x: s.snake.body[0].x - 1, y: s.snake.body[0].y}
			if leftCoord != s.snake.body[1] {
				s.snake.direction = LEFT
			}
		}
	case TickMsg:
		// move snake head in direction
		prevSnakePartPos := s.snake.body[0]

		switch s.snake.direction {
		case UP:
			s.snake.body[0].y -= 1
		case RIGHT:
			s.snake.body[0].x += 1
		case DOWN:
			s.snake.body[0].y += 1
		case LEFT:
			s.snake.body[0].x -= 1
		}

		if s.snake.body[0].x < 0 || s.snake.body[0].x >= BOARD_WIDTH ||
			s.snake.body[0].y < 0 || s.snake.body[0].y >= BOARD_HEIGHT ||
			s.snake.isHeadCollidingWithBody() {

			s.gameOver = true
			break
		}

		atePellet := s.snake.body[0] == s.pellet

		/*
			move the rest of the snake
			temporarily save position of current part as prevPos
			move part to prevSnakePartPos
			set prevSnakePartPos to prevPos for the next iteration
		*/
		for i := 1; i < len(s.snake.body); i++ {
			prevPos := s.snake.body[i]
			s.snake.body[i] = prevSnakePartPos
			prevSnakePartPos = prevPos
		}

		if atePellet {
			s.snake.body = append(s.snake.body, prevSnakePartPos)
			s.spawnPellet()
			s.score++
		}

		s.updateBoard()

		return s, tickEvery()
	}

	return s, nil
}
