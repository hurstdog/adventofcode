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

const EDGE = 1000
const NUM_LIGHTS = EDGE * EDGE

// Grid[x][y] == brightness
type Grid [EDGE][EDGE]int

// Fun fact: initialized to all off
var lights Grid

func InitLights() {
	for i := 0; i < EDGE; i++ {
		for j := 0; j < EDGE; j++ {
			setLight(Point{i, j}, OFF)
		}
	}
}

// Given a point, returns the brightness stored at that Point in the grid, or
// an error.
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

// Applies the given Command to the stored light grid
func ApplyCmd(c Command) error {
	for i := c.start.x; i <= c.end.x; i++ {
		for j := c.start.y; j <= c.end.y; j++ {
			err := setLight(Point{i, j}, c.cmd)
			if err != nil {
				return fmt.Errorf("ApplyCmd: %v\n", err)
			}
		}
	}
	return nil
}

// Returns the number of lights set to ON
func NumOn() int {
	var c int
	for i := 0; i < EDGE; i++ {
		for j := 0; j < EDGE; j++ {
			c += lights[i][j]
		}
	}

	return c
}

// LineToCmd takes a string from the input file and returns a command struct
// containing the instruction to execute.
// Lines must match the format: "[turn off|turn on|toggle] N,N through N,N"
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
