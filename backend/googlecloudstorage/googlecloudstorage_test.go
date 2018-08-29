// Test GoogleCloudStorage filesystem interface
package googlecloudstorage_test

import (
	"testing"

	"github.com/adbegon/rclone/backend/googlecloudstorage"
	"github.com/adbegon/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestGoogleCloudStorage:",
		NilObject:  (*googlecloudstorage.Object)(nil),
	})
}
