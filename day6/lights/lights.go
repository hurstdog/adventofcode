// Lights is a package for processing the light map from day6
package lights

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	OFF = iota
	ON
	TOGGLE
)

type Point struct {
	x, y  int
	value int // ON|OFF
}

type Command struct {
	cmd   int
	start Point
	end   Point
}

const EDGE = 1000

// Grid[x][y] == Point
type Grid [EDGE][EDGE]int

var lights Grid

// Given a point, returns the value stored at that Point in the grid, or an
// error.
func getLight(p Point) (int, error) {
	if p.x >= EDGE || p.y >= EDGE {
		return OFF, fmt.Errorf("Point %v out of range %d\n", p, EDGE)
	}

	return lights[p.x][p.y], nil
}

// setLight takes a command of ON|OFF|TOGGLE and a Point, and affects that
// instruction on that point in the Grid. This DOES NOT modify the incoming
// Point.
func setLight(p Point, cmd int) error {
	if p.x >= EDGE || p.y >= EDGE {
		return fmt.Errorf("Point %v out of range %d\n", p, EDGE)
	}
	switch cmd {
	case OFF:
		lights[p.x][p.y] = OFF
	case ON:
		lights[p.x][p.y] = ON
	case TOGGLE:
		if lights[p.x][p.y] == OFF {
			lights[p.x][p.y] = ON
		} else {
			lights[p.x][p.y] = OFF
		}
	}

	return nil
}

// LineToCmd takes a string from the input file and returns a command struct
// containing the instruction to execute.
// lines must match the format: "[turn off|turn on|toggle] N,N through N,N"
func LineToCmd(line string) (Command, error) {
	var offset int
	var c Command
	t := strings.Split(line, " ")
	if t[0] == "toggle" {
		c.cmd = TOGGLE
	} else if t[0] == "turn" {
		if t[1] == "on" {
			c.cmd = ON
		} else {
			c.cmd = OFF
		}
		offset = 1
	} else {
		return Command{}, fmt.Errorf("Unknown command %s\n", t[0])
	}

	var err error
	c.start, err = parsePoint(t[1+offset])
	c.end, err = parsePoint(t[3+offset])

	return c, err
}

// parsePoint parses text like "N,M" and returns a Point{N, M}.
func parsePoint(point string) (Point, error) {
	parts := strings.Split(point, ",")
	if len(parts) != 2 {
		return Point{}, fmt.Errorf("%s doesn't match n,m\n", point)
	}
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return Point{}, err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return Point{}, err
	}

	return Point{x, y, OFF}, nil
}
