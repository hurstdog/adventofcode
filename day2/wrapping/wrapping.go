// wrapping is a library for use in wrapping paper calculations
package wrapping

import (
	"fmt"
	"strconv"
	"strings"
)

// RibbonNeeded calculates the ribbon needed to wrap a box of size LxWxH.
// Len ribbon is the perimeter of the smallest face plus the volume of the box.
// So if LxW is the smallest face, we need 2*(L+W) + L*W*H
func RibbonNeeded(s string) (int, error) {
	lwh, err := toLWH(s)
	if err != nil {
		return -1, err
	}

	// Handy shorthand
	l := lwh[0]
	w := lwh[1]
	h := lwh[2]

	// Perimeters
	lw := 2 * (l + w)
	lh := 2 * (l + h)
	wh := 2 * (w + h)

	// Find the smallest perimeter
	least := lw
	if lh < least {
		least = lh
	}
	if wh < least {
		least = wh
	}

	vol := l * w * h
	return least + vol, nil
}

// PaperNeeded calculates the amount of wrapping paper needed given length,
// width, height of a box. LWH needs to be given in a string like "LxWxH".
// Amount needed is 2lw + 2lh + 2hw + the area of the smallest side.
func PaperNeeded(s string) (int, error) {
	lwh, err := toLWH(s)
	if err != nil {
		return -1, err
	}

	// Handy shorthand
	l := lwh[0]
	w := lwh[1]
	h := lwh[2]

	lw := l * w
	lh := l * h
	wh := w * h

	// Find the smallest area
	least := lw
	if lh < least {
		least = lh
	}
	if wh < least {
		least = wh
	}

	return 2*lw + 2*lh + 2*wh + least, nil
}

// toLWH takes a string like NxMxO and returns N, M, and O separately.
func toLWH(s string) ([]int, error) {
	var err error
	var res []int = make([]int, 3)
	dems := strings.Split(s, "x")
	if len(dems) != 3 {
		return nil, fmt.Errorf("%s doesn't split to 3.\n", s)
	}
	for i, v := range dems {
		res[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
