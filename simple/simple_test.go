package simple

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"os"
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
	require.Greater(t, rand, 10, "This test flaked. It should flake ten percent of the time")

	t.Log("This test is healthy. It should flake ten percent of the time")
}

func TestFlakyTwentyFivePercent(t *testing.T) {
	t.Parallel()

	rand := rand.Intn(100)
	require.Greater(t, rand, 25, "This test flaked. It should flake twenty five percent of the time")

	t.Log("This test is healthy. It should flake twenty five percent of the time")
}

func TestFlakyFortyNinePercent(t *testing.T) {
	t.Parallel()

	rand := rand.Intn(100)
	require.Greater(t, rand, 49, "This test flaked. It should flake 49 percent of the time")

	t.Log("This test is healthy. It should flake 49 percent of the time")
}

func TestFlakyFiftyPercent(t *testing.T) {
	if os.Getenv("RUN_QUARANTINED_TESTS") != "true" {
		t.Skip("Flaky test quarantined. Ticket <Jira ticket>. Done automatically by branch-out (https://github.com/smartcontractkit/branch-out)")
	} else {
		t.Logf("'RUN_QUARANTINED_TESTS' set to '%s', running quarantined test", os.Getenv("RUN_QUARANTINED_TESTS"))
	}
	t.Parallel()

	rand := rand.Intn(100)
	require.Greater(t, rand, 50, "This test flaked. It should flake fifty percent of the time")

	t.Log("This test is healthy. It should flake fifty percent of the time")
}

func TestFlakySeventyFivePercent(t *testing.T) {
	t.Parallel()

	rand := rand.Intn(100)
	require.Greater(t, rand, 75, "This test flaked. It should flake seventy five percent of the time")

	t.Log("This test is healthy. It should flake seventy five percent of the time")
}

func TestFlakyNinetyPercent(t *testing.T) {
	t.Parallel()

	rand := rand.Intn(100)
	require.Greater(t, rand, 90, "This test flaked. It should flake ninety percent of the time")

	t.Log("This test is healthy. It should flake ninety percent of the time")
}
