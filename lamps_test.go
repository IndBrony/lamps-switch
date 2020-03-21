package lamps

import "testing"

func TestLampSwitch(t *testing.T) {
	basicTesting := func(t *testing.T, lampsNum, timesPressed, expected int) {
		if output := LampSwitch(lampsNum, timesPressed); output != expected {
			t.Errorf(`
Test Failed With Input :
	lampsNum		: %v
	timesPressed	: %v
Expecting %v but got %v`,
				lampsNum, timesPressed, expected, output,
			)
		}
	}
	basicTesting(t, 10, 1, 10)
	basicTesting(t, 10, 2, 5)
	basicTesting(t, 10, 3, 4)
	basicTesting(t, 10, 4, 6)
}
