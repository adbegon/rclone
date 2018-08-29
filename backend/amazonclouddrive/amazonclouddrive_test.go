// Test AmazonCloudDrive filesystem interface

// +build acd

package amazonclouddrive_test

import (
	"testing"

	"github.com/adbegon/rclone/backend/amazonclouddrive"
	"github.com/adbegon/rclone/fs"
	"github.com/adbegon/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.NilObject = fs.Object((*amazonclouddrive.Object)(nil))
	fstests.RemoteName = "TestAmazonCloudDrive:"
	fstests.Run(t)
}
