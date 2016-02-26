package arukas

import (
	"testing"
)

func ExampleStart() {
	runCommand([]string{"arukas", "start", "2b21fe34-328f-4d7e-8678-726d9eff2b7f"})
	// Output:
	// Starting...
}

func TestStartAlreadyRunning(t *testing.T) {
	runCommand([]string{"arukas", "start", "d19b004c-0d59-4f4f-955c-5bace7c49a34"})
	if ExitCode != 1 {
		t.Errorf(("ExitCode got: %d, want: 1"), ExitCode)
	}
}

func ExampleStartAlreadyRunning() {
	runCommand([]string{"arukas", "start", "d19b004c-0d59-4f4f-955c-5bace7c49a34"})
	// Output:
	// Failed to start the container
}
