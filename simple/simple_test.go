package simple

import (
	"math/rand"
	"testing"
)

func TestHealthy(t *testing.T) {
	t.Parallel()
	t.Log("This test is healthy")
}

func TestBroken(t *testing.T) {
	t.Parallel()
	t.Fatal("This test is intentionally broken")
}

func TestSkip(t *testing.T) {
	t.Parallel()

	t.Skip("This test is intentionally skipped")
}

func TestFlakyTenPercent(t *testing.T) {
	t.Parallel()

	if rand.Intn(10) < 1 {
		t.Fatal("This test flaked. It should flake ten percent of the time")
	}

	t.Log("This test is healthy. It should flake ten percent of the time")
}

func TestFlakyTwentyFivePercent(t *testing.T) {
	t.Parallel()

	if rand.Intn(100) < 25 {
		t.Fatal("This test flaked. It should flake twenty five percent of the time")
	}

	t.Log("This test is healthy. It should flake twenty five percent of the time")
}

func TestFlakyFiftyPercent(t *testing.T) {
	t.Parallel()

	if rand.Intn(100) < 50 {
		t.Fatal("This test flaked. It should flake fifty percent of the time")
	}

	t.Log("This test is healthy. It should flake fifty percent of the time")
}

func TestFlakySeventyFivePercent(t *testing.T) {
	t.Parallel()

	if rand.Intn(100) < 75 {
		t.Fatal("This test flaked. It should flake seventy five percent of the time")
	}

	t.Log("This test is healthy. It should flake seventy five percent of the time")
}
