package calculator

import (
	"math"
	"testing"
)

func TestCalculator(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("0.3 plus 0.2 equals 0.5", func(t *testing.T) {
		got := roundToTwoDecimalPoints(Sum(0.3, 0.2))
		want := 0.5
		assertCorrectMessage(t, got, want)
	})

	t.Run("3.5 plus 2.5 equals 6.0", func(t *testing.T) {
		got := roundToTwoDecimalPoints(Sum(3.5, 2.5))
		want := 6.0
		assertCorrectMessage(t, got, want)
	})

	t.Run("0.2 plus 0.04 equals 0.24", func(t *testing.T) {
		got := roundToTwoDecimalPoints(Sum(0.2, 0.04))
		want := 0.24
		assertCorrectMessage(t, got, want)
	})

	t.Run("0.36 plus 0.04 equals 0.4", func(t *testing.T) {
		got := roundToTwoDecimalPoints(Sum(0.36, 0.04))
		want := 0.4
		assertCorrectMessage(t, got, want)
	})

	t.Run("0.68 plus 0.04 equals 0.72", func(t *testing.T) {
		got := roundToTwoDecimalPoints(Sum(0.68, 0.04))
		want := 0.72
		assertCorrectMessage(t, got, want)
	})
}

func roundToTwoDecimalPoints(value float64) float64 {
	return math.Round(value*100) / 100
}
