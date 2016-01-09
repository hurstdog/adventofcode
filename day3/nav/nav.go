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
