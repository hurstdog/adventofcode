// nav is a library for interpreting directions like <>^v
package nav

import (
	"fmt"
	"strings"
)

type Point struct {
	x int
	y int
}

var pointsSeen map[Point]int = make(map[Point]int)

// Given a string of directions like <>^v... this will run the directions and
// return a count of the number of points that get visited at least once.
func AtLeastOne(dirs string) (int, error) {
	ResetPoints()
	err := handleDirections(dirs)
	if err != nil {
		return -1, err
	}

	return len(pointsSeen), nil
}

// Same as at least one, but using a robo santa too.
func AtLeastOneRobo(dirs string) (int, error) {
	ResetPoints()
	err := handleRoboSantaDirections(dirs)
	if err != nil {
		return -1, err
	}

	return len(pointsSeen), nil
}

func ResetPoints() {
	for k, _ := range pointsSeen {
		delete(pointsSeen, k)
	}
}

func handleDirections(dirs string) error {
	// Inialize the first point
	p := Point{0, 0}
	pointsSeen[p] = 1
	for _, d := range strings.Split(dirs, "") {
		var err error
		p, err = updateLoc(p, d)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		pointsSeen[p]++
	}

	return nil
}

// Same as HandleDirections, but this works with robosanta, where half the
// directions go to one santa, half to the other.
func handleRoboSantaDirections(dirs string) error {
	// Inialize the first point, used for both santas
	p := Point{0, 0}
	pointsSeen[p] = 1
	rp := p // rp == robo point
	for i, d := range strings.Split(dirs, "") {
		var err error
		if i%2 == 0 {
			p, err = updateLoc(p, d)
			pointsSeen[p]++
		} else {
			rp, err = updateLoc(rp, d)
			pointsSeen[rp]++
		}
		if err != nil {
			return fmt.Errorf("%v", err)
		}
	}

	return nil
}

// updateLoc takes a coordinate and a direction, and returns the new
// coordinate.
func updateLoc(p Point, d string) (Point, error) {
	switch d {
	case ">":
		p.x++
	case "<":
		p.x--
	case "^":
		p.y++
	case "v":
		p.y--
	default:
		err := fmt.Errorf("'%s' not a supported directino")
		return Point{}, err
	}
	return p, nil
}
