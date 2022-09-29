package upgrade

import (
	"fmt"

	"github.com/leoay/luna/cmd/luna/v2/internal/base"

	"github.com/spf13/cobra"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the luna tools",
	Long:  "Upgrade the luna tools. Example: luna upgrade",
	Run:   Run,
}

// Run upgrade the kratos tools.
func Run(cmd *cobra.Command, args []string) {
	err := base.GoInstall(
		"github.com/leoay/luna/cmd/luna/v2@latest",
	)
	if err != nil {
		fmt.Println(err)
	}
}
