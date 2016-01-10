// nice provides methods to determine if a given string is naughty or nice.
package nice

import (
	"strings"
)

// Returns true if a string is nice, false, if it's naughty
func Nice(s string) bool {
	return repeats(s) && vowels(s) && !banned(s)
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
