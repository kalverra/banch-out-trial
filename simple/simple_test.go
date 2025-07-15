package simple

import (
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// runFailingTests checks if failing tests should be run based on environment variable
func runFailingTests() bool {
	return os.Getenv("RUN_FAILING_TESTS") == "true"
}

func TestHealthy(t *testing.T) {
	t.Parallel()
	t.Log("This test is healthy")
}

func TestBroken(t *testing.T) {
	t.Parallel()

	if runFailingTests() {
		require.Fail(t, "This test is intentionally broken")
	}
	t.Log("This test is healthy when flaky tests are disabled")
}

func TestSkip(t *testing.T) {
	t.Parallel()

	t.Skip("This test is intentionally skipped")
}

func TestFlakyTenPercent(t *testing.T) {
	t.Parallel()

	if runFailingTests() {
		rand := rand.Intn(100)
		require.Greater(t, rand, 10, "This test flaked. It should flake ten percent of the time")
	}
	t.Log("This test is healthy. It should flake ten percent of the time unless RUN_FAILING_TESTS is set to false")
}

func TestFlakyTwentyFivePercent(t *testing.T) {
	t.Parallel()

	if runFailingTests() {
		rand := rand.Intn(100)
		require.Greater(t, rand, 25, "This test flaked. It should flake twenty five percent of the time")
	}
	t.Log("This test is healthy. It should flake twenty five percent of the time unless RUN_FAILING_TESTS is set to false")
}

func TestFlakyFortyNinePercent(t *testing.T) {
	t.Parallel()

	if runFailingTests() {
		rand := rand.Intn(100)
		require.Greater(t, rand, 49, "This test flaked. It should flake 49 percent of the time")
	}
	t.Log("This test is healthy. It should flake 49 percent of the time unless RUN_FAILING_TESTS is set to false")
}

func TestFlakyFiftyPercent(t *testing.T) {
	t.Parallel()

	if runFailingTests() {
		rand := rand.Intn(100)
		require.Greater(t, rand, 50, "This test flaked. It should flake fifty percent of the time")
	}
	t.Log("This test is healthy. It should flake fifty percent of the time unless RUN_FAILING_TESTS is set to false")
}

func TestFlakySeventyFivePercent(t *testing.T) {
	t.Parallel()

	if runFailingTests() {
		rand := rand.Intn(100)
		require.Greater(t, rand, 75, "This test flaked. It should flake seventy five percent of the time")
	}
	t.Log("This test is healthy. It should flake seventy five percent of the time unless RUN_FAILING_TESTS is set to false")
}

func TestFlakyNinetyPercent(t *testing.T) {
	t.Parallel()

	if runFailingTests() {
		rand := rand.Intn(100)
		require.Greater(t, rand, 90, "This test flaked. It should flake ninety percent of the time")
	}
	t.Log("This test is healthy. It should flake ninety percent of the time unless RUN_FAILING_TESTS is set to false")
}
