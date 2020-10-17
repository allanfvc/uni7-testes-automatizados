package calculator

import (
	"math"
	"testing"
)

func TestCalculator(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want float64) {
		t.Helper()
		if(got != want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("0.2 plus 0.04 equals 0.24", func(t *testing.T) {
		got := roundToTwoDecimalPoints(Sum(0.2, 0.04))
		want := 0.24
		assertCorrectMessage(t, got, want)
	})

	t.Run("30.0 minus 20.0 equals 10.0", func(t *testing.T) {
		got := roundToTwoDecimalPoints(Subtract(30.0, 20.0))
		want := 10.00
		assertCorrectMessage(t, got, want)
	})
}

func roundToTwoDecimalPoints(value float64) float64 {
	return math.Round(value * 100) /100;
}