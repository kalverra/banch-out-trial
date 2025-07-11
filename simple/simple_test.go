package simple

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealthy(t *testing.T) {
	t.Parallel()
	t.Log("This test is healthy")
}

func TestBroken(t *testing.T) {
	t.Parallel()

	require.Fail(t, "This test is intentionally broken")
}

func TestSkip(t *testing.T) {
	t.Parallel()

	t.Skip("This test is intentionally skipped")
}

func TestFlakyTenPercent(t *testing.T) {
	t.Parallel()

	rand := rand.Intn(100)
	require.Less(t, rand, 10, "This test flaked. It should flake ten percent of the time")

	t.Log("This test is healthy. It should flake ten percent of the time")
}

func TestFlakyTwentyFivePercent(t *testing.T) {
	t.Parallel()

	rand := rand.Intn(100)
	require.Less(t, rand, 25, "This test flaked. It should flake twenty five percent of the time")

	t.Log("This test is healthy. It should flake twenty five percent of the time")
}

func TestFlakyFiftyPercent(t *testing.T) {
	t.Parallel()

	rand := rand.Intn(100)
	require.Less(t, rand, 50, "This test flaked. It should flake fifty percent of the time")

	t.Log("This test is healthy. It should flake fifty percent of the time")
}

func TestFlakySeventyFivePercent(t *testing.T) {
	t.Parallel()

	rand := rand.Intn(100)
	require.Less(t, rand, 75, "This test flaked. It should flake seventy five percent of the time")

	t.Log("This test is healthy. It should flake seventy five percent of the time")
}
