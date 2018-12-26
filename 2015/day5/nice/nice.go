// nice provides methods to determine if a given string is naughty or nice.
package nice

import (
	"strings"
)

// Returns true if a string is nice, false, if it's naughty
func Nice(s string) bool {
	return repeats(s) && vowels(s) && !banned(s)
}

// Nice2 returns true using doubledouble and gaprepeat
func Nice2(s string) bool {
	return doubledouble(s) && gaprepeat(s)
}

// doubledouble returns true if there is a two letter string that repeats
// itself twice, without overlapping.
func doubledouble(s string) bool {
	max := len(s)
	if max < 4 {
		return false
	}
	for i, _ := range s {
		// don't overflow
		if i > max-3 {
			break
		}
		match := strings.Index(s[i+2:], s[i:i+2])
		if match >= 0 {
			return true
		}
	}
	return false
}

// gaprepeat returns true if the string has the pattern xyx, where x is any
// character.
func gaprepeat(s string) bool {
	max := len(s)
	if max < 3 {
		return false
	}
	for i, _ := range s {
		if i+2 >= max {
			break
		}
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

// repeats returns true if the string has a character that repeats itself.
func repeats(s string) bool {
	if len(s) < 2 {
		return false
	}
	last := s[0]
	for i := 1; i < len(s); i++ {
		if last == s[i] {
			return true
		}
		last = s[i]
	}

	return false
}

// vowels returns true if the string has at least three distinct vowels.
func vowels(s string) bool {
	var c int
	for _, v := range []string{"a", "e", "i", "o", "u"} {
		c += strings.Count(s, v)
	}

	return c >= 3
}

// banned returns true if the string contains banned character combinations
func banned(s string) bool {
	for _, v := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Index(s, v) >= 0 {
			return true
		}
	}
	return false
}
