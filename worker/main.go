package main

import (
	"image"
	"log"
	"time"
)

type command int

const (
	right = command(0)
	left = command(1)
)

type RoverDriver struct {
	commandc chan command
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}

	r.drive()
	return r
}

// drive is responsible for driving the rover. It
// is expected to be started in a goroutine
func (r *RoverDriver) drive() {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <- r.commandc:
			switch c {
			case right:
				direction = image.Point{X: -direction.Y, Y: direction.X}
			case left:
				direction = image.Point{X: direction.Y, Y: -direction.X}
			}
			log.Printf("new direction %v", direction)
		case <- nextMove:
			pos = pos.Add(direction)
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

// Left turns the rover left (90 degress counterclockwise)
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// Right turns the rover right (90 degress clockwise)
func (r *RoverDriver) Right() {
	r.commandc <- right
}

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}