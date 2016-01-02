// parens is a library that counts parens and returns a number
package parens

// count reads through a string and adds 1 for every (, -1 for every ),
// returning the total count.
func Count(s []byte) int {
	var c int
	for _, paren := range s {
		switch paren {
		case ')':
			c--
		case '(':
			c++
		}
	}
	return c
}

// Like Count(), but instead of returning the floor Santa ends up on this
// returns the first floor where we go negative.
func Position(s []byte) int {
	var c int
	for i, paren := range s {
		switch paren {
		case ')':
			c--
			if c < 0 {
				return i + 1
			}
		case '(':
			c++
		}
	}
	// Doesn't mean anything to get here.
	return -1
}
