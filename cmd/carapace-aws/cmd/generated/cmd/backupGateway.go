package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGatewayCmd = &cobra.Command{
	Use:   "backup-gateway",
	Short: "Backup gateway",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupGatewayCmd).Standalone()

	})
	rootCmd.AddCommand(backupGatewayCmd)
}
