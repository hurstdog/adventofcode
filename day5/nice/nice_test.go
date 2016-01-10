// Tests for the nice package
package nice

import (
	"testing"
)

func TestNice(t *testing.T) {
	testNice("ugknbfddgicrmopn", true, t)
	testNice("aaa", true, t)
	testNice("jchzalrnumimnmhp", false, t) // no double letter
	testNice("haegwjzuvuyypxyu", false, t) // contains xy
	testNice("dvszwmarrgswjxmb", false, t) // one vowel
}

func testNice(s string, exp bool, t *testing.T) {
	res := Nice(s)
	if res != exp {
		t.Errorf("Error: Nice(%s) == %t, expected %t", s, res, exp)
	}
}

func TestHasRepeat(t *testing.T) {
	testRepeat("abcde", false, t)
	testRepeat("abcdd", true, t)
	testRepeat("aabcd", true, t)
	testRepeat("abababab", false, t)
	testRepeat("", false, t)
	testRepeat("a", false, t)
}

func testRepeat(s string, exp bool, t *testing.T) {
	res := repeats(s)
	if res != exp {
		t.Errorf("Error: repeats(%s) == %v, expected %v", s, res, exp)
	}
}

func TestVowels(t *testing.T) {
	testVowels("", false, t)
	testVowels("aei", true, t)
	testVowels("bou", false, t)
	testVowels("aeiou", true, t)
	testVowels("abicdafghoa", true, t)
	testVowels("aaa", true, t)
	testVowels("aa", false, t)
}

func testVowels(s string, exp bool, t *testing.T) {
	res := vowels(s)
	if res != exp {
		t.Errorf("Error: vowels(%s) == %v, expected %v", s, res, exp)
	}
}

func TestBanned(t *testing.T) {
	testBanned("", false, t)
	testBanned("a", false, t)
	testBanned("ab", true, t)
	testBanned("cd", true, t)
	testBanned("pq", true, t)
	testBanned("xy", true, t)
	testBanned("acb", false, t)
}

func testBanned(s string, exp bool, t *testing.T) {
	res := banned(s)
	if res != exp {
		t.Errorf("Error: banned(%s) == %v, expected %v", s, res, exp)
	}
}
