package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_updateServerConfigCmd = &cobra.Command{
	Use:   "update-server-config",
	Short: "Updates the configuration of the specified server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_updateServerConfigCmd).Standalone()

	migrationhubstrategy_updateServerConfigCmd.Flags().String("server-id", "", "The ID of the server.")
	migrationhubstrategy_updateServerConfigCmd.Flags().String("strategy-option", "", "The preferred strategy options for the application component.")
	migrationhubstrategy_updateServerConfigCmd.MarkFlagRequired("server-id")
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_updateServerConfigCmd)
}
