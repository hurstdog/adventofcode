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
	x, y int
}

type Command struct {
	cmd   int
	start Point
	end   Point
}

// Grid[x][y] == [ON|OFF]
type Grid map[int]map[int]int

var lights Grid = make(Grid)

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

	return Point{x, y}, nil
}
