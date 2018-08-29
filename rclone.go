// Sync files and directories to and from local and remote object stores
//
// Nick Craig-Wood <nick@craig-wood.com>
package main

import (
	"github.com/adbegon/rclone/cmd"

	_ "github.com/adbegon/rclone/backend/all" // import all backends
	_ "github.com/adbegon/rclone/cmd/all"     // import all commands
)

func main() {
	cmd.Main()
}
