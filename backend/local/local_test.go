// Test Local filesystem interface
package local_test

import (
	"testing"

	"github.com/adbegon/rclone/backend/local"
	"github.com/adbegon/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "",
		NilObject:  (*local.Object)(nil),
	})
}
