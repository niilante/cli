package arukas

import (
	"testing"
)

func TestRemove(t *testing.T) {
	runCommand([]string{"arukas", "rm", "2b21fe34-328f-4d7e-8678-726d9eff2b7f"})
	if ExitCode != 0 {
		t.Errorf(("ExitCode got: %d, want: 0"), ExitCode)
	}
}
