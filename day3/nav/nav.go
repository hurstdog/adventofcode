// nav is a library for interpreting directions like <>^v
package nav

import (
	"fmt"
)

type Point struct {
	x int
	y int
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
