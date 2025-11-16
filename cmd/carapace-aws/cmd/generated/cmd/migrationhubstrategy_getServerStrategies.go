package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_getServerStrategiesCmd = &cobra.Command{
	Use:   "get-server-strategies",
	Short: "Retrieves recommended strategies and tools for the specified server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_getServerStrategiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubstrategy_getServerStrategiesCmd).Standalone()

		migrationhubstrategy_getServerStrategiesCmd.Flags().String("server-id", "", "The ID of the server.")
		migrationhubstrategy_getServerStrategiesCmd.MarkFlagRequired("server-id")
	})
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_getServerStrategiesCmd)
}
