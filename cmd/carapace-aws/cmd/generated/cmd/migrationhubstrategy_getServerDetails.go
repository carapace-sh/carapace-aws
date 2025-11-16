package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_getServerDetailsCmd = &cobra.Command{
	Use:   "get-server-details",
	Short: "Retrieves detailed information about a specified server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_getServerDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubstrategy_getServerDetailsCmd).Standalone()

		migrationhubstrategy_getServerDetailsCmd.Flags().String("max-results", "", "The maximum number of items to include in the response.")
		migrationhubstrategy_getServerDetailsCmd.Flags().String("next-token", "", "The token from a previous call that you use to retrieve the next set of results.")
		migrationhubstrategy_getServerDetailsCmd.Flags().String("server-id", "", "The ID of the server.")
		migrationhubstrategy_getServerDetailsCmd.MarkFlagRequired("server-id")
	})
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_getServerDetailsCmd)
}
