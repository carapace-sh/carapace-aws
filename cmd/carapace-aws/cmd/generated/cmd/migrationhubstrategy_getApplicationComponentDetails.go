package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_getApplicationComponentDetailsCmd = &cobra.Command{
	Use:   "get-application-component-details",
	Short: "Retrieves details about an application component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_getApplicationComponentDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubstrategy_getApplicationComponentDetailsCmd).Standalone()

		migrationhubstrategy_getApplicationComponentDetailsCmd.Flags().String("application-component-id", "", "The ID of the application component.")
		migrationhubstrategy_getApplicationComponentDetailsCmd.MarkFlagRequired("application-component-id")
	})
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_getApplicationComponentDetailsCmd)
}
