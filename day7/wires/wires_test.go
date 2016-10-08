// Tests for the Wires package
package wires

import (
	"fmt"
	"testing"
)

func TestRunLine(t *testing.T) {
	expectValue("123 -> x", "x", 123, t)
	exp := fmt.Errorf("Line [123 -> x] assigns to a token that already exists.\n")
	expectErr("123 -> x", exp, t)
}

func expectValue(s string, k string, v int, t *testing.T) {
	err := RunLine(s)
	if err != nil {
		t.Error(err)
	}
	if C[k] != v {
		t.Errorf("Expecting %v = %v, got %v\n", k, v, C[k])
	}
}

func expectErr(s string, e error, t *testing.T) {
	err := RunLine(s)
	if err.Error() != e.Error() {
		t.Errorf("Expected error [%v], got error [%v]\n", e, err)
	}
}

/*
var testData map[string]Command = make(map[string]Command)

func init() {
	testData["turn on 0,0 through 999,999"] =
		Command{ON, Point{0, 0}, Point{999, 999}}
	testData["turn off 0,0 through 999,999"] =
		Command{OFF, Point{0, 0}, Point{999, 999}}
	testData["toggle 0,0 through 999,999"] =
		Command{TOGGLE, Point{0, 0}, Point{999, 999}}
	testData["toggle 888,888 through 888,888"] =
		Command{TOGGLE, Point{888, 888}, Point{888, 888}}
}

func TestSetLight(t *testing.T) {
	p := Point{0, 0}
	v, err := getLight(p)
	expectValue(v, 0, err, t)

	// Can't go below 0
	err = setLight(p, OFF)
	v, err = getLight(p)
	expectValue(v, 0, err, t)

	// OFF -> ON
	err = setLight(p, ON)
	v, err = getLight(p)
	expectValue(v, ON, err, t)

	// Increase brightness by 1
	err = setLight(p, ON)
	v, err = getLight(p)
	expectValue(v, 2, err, t)

	// Toggle is +2
	err = setLight(p, TOGGLE)
	v, err = getLight(p)
	expectValue(v, 4, err, t)

	// OFF is decrease by 1
	err = setLight(p, OFF)
	v, err = getLight(p)
	expectValue(v, 3, err, t)

	// -1 is reset
	err = setLight(p, -1)
	v, err = getLight(p)
	expectValue(v, 0, err, t)
}

func TestTotBrightness(t *testing.T) {
	ResetLights()
	n := TotBrightness()
	if n != 0 {
		t.Errorf("Expected 0 lights on, got %d\n", n)
	}
	ApplyCmd(Command{ON, Point{0, 0}, Point{EDGE - 1, EDGE - 1}})
	n = TotBrightness()
	if n != NUM_LIGHTS {
		t.Errorf("Expected %d lights on, got %d\n", NUM_LIGHTS, n)
	}
}
*/
