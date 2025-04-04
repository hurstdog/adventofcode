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

func TestDoubleDouble(t *testing.T) {
	testDoubleDouble("xyxy", true, t)
	testDoubleDouble("aabcdefgaa", true, t)
	testDoubleDouble("aaa", false, t)
}

func testDoubleDouble(s string, exp bool, t *testing.T) {
	res := doubledouble(s)
	if res != exp {
		t.Errorf("Error: doubledouble(%s) == %v, expected %v", s, res, exp)
	}
}

func TestGapRepeat(t *testing.T) {
	testGapRepeat("", false, t)
	testGapRepeat("a", false, t)
	testGapRepeat("aa", false, t)
	testGapRepeat("aaa", true, t)
	testGapRepeat("xyx", true, t)
	testGapRepeat("aabcdefe", true, t)
	testGapRepeat("aabcdefg", false, t)
}

func testGapRepeat(s string, exp bool, t *testing.T) {
	res := gaprepeat(s)
	if res != exp {
		t.Errorf("Error: doubledouble(%s) == %v, expected %v", s, res, exp)
	}
}

func TestNice2(t *testing.T) {
	testNice2("qjhvhtzxzqqjkmpb", true, t)
	testNice2("xxyxx", true, t)
	testNice2("uurcxstgmygtbstg", false, t)
	testNice2("ieodomkazucvgmuy", false, t)
}

func testNice2(s string, exp bool, t *testing.T) {
	res := Nice2(s)
	if res != exp {
		t.Errorf("Error: Nice2(%s) == %t, expected %t", s, res, exp)
	}
}
