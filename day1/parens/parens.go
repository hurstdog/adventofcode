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
