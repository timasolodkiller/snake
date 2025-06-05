package game

import (
	"game_snake/game_board"
	"game_snake/render"
	"game_snake/snake"
	"github.com/nsf/termbox-go"
	"game_snake/point"
	"time"
	"math/rand"
	"game_snake/food"
	"os"
)
var (
	Snake *snake.Snake
	Board *board.Board
	eventChan = make(chan termbox.Event)
	foodChan = make(chan bool, 1)
	random *rand.Rand
	Food food.Food
	Logger *os.File
)

func Init() {
	Board = board.Initialize(50, 30)
	Snake = snake.Initialize(point.Point{X : Board.Height / 2, Y : Board.Width / 2}, snake.Right, Board)
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
	Logger, _ = os.Create("log.txt")
	Food = food.Initialize(Snake, random, Logger)
}



func ListenerForKeyBoard() {
	for {
		eventChan <- termbox.PollEvent();
	}
}
func Controller(s *snake.Snake) snake.Direction {
	select {
	case v := <- eventChan:
			if v.Type == termbox.EventKey {
				switch v.Key {
				case termbox.KeyArrowUp:
					return snake.Up
				case termbox.KeyArrowRight:
					return snake.Right
				case termbox.KeyArrowLeft:
					return snake.Left
				case termbox.KeyArrowDown:
					return snake.Down
				}
			}
		default:
			return s.GetHeadDirection()
		}
	return s.GetHeadDirection()
}
func ListenerForFood() {
	for {
		answer, ok := <- foodChan
		if answer {
			Food = food.Initialize(Snake, random, Logger)
		}
		if !ok {
			break
		}
	}
}

func Start() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	go ListenerForKeyBoard()
	go ListenerForFood()
	defer close(foodChan)
	defer close(eventChan)
	defer termbox.Close()
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		render.Render(Snake, Board, Food)
		termbox.Flush()
		direction := Controller(Snake)
		if !Snake.Move(direction) && !Board.IsCrashed(Snake.GetHeadCoord()) {
			break
		}
		if Snake.GetHeadCoord() == Food.Coord {
			Snake.Grow()
			foodChan <- true
		}
		time.Sleep(100 * time.Millisecond)
	}
}