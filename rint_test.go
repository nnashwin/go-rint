package rint_test

import (
	"github.com/ru-lai/rint"
	"testing"
)

func checkInRange(num, start, end int) bool {
	return num >= start && num <= end
}

func TestGetRint(t *testing.T) {
	testLimit := 10
	actual := rint.GetRint(testLimit)
	if !checkInRange(rint.GetRint(testLimit), 0, 9) {
		t.Errorf("The GetRint test has generated an illegal value outside of the range 0, 9.  testLimit: %d, actual: %d", testLimit, actual)
	}
}

func TestGetRintRange(t *testing.T) {
	cases := []struct {
		tMin, tMax int
	}{
		{
			tMin: 5,
			tMax: 6,
		},
		{
			tMin: 0,
			tMax: 10,
		},
		{
			tMin: 4,
			tMax: 10,
		},
		{
			tMin: 4000,
			tMax: 5000,
		},
	}

	for _, c := range cases {
		if !checkInRange(rint.GetRintRange(c.tMin, c.tMax), c.tMin, c.tMax) {
			t.Errorf("The GetRintRange test has generated a value outside of the range. Range: Min: %d, Max: %d", c.tMin, c.tMax)
		}
	}
}
