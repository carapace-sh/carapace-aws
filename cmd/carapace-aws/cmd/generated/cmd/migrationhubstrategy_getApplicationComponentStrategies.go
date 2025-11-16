package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_getApplicationComponentStrategiesCmd = &cobra.Command{
	Use:   "get-application-component-strategies",
	Short: "Retrieves a list of all the recommended strategies and tools for an application component running on a server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_getApplicationComponentStrategiesCmd).Standalone()

	migrationhubstrategy_getApplicationComponentStrategiesCmd.Flags().String("application-component-id", "", "The ID of the application component.")
	migrationhubstrategy_getApplicationComponentStrategiesCmd.MarkFlagRequired("application-component-id")
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_getApplicationComponentStrategiesCmd)
}
