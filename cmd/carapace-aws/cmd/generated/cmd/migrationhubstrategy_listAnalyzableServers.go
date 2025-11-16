package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_listAnalyzableServersCmd = &cobra.Command{
	Use:   "list-analyzable-servers",
	Short: "Retrieves a list of all the servers fetched from customer vCenter using Strategy Recommendation Collector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_listAnalyzableServersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubstrategy_listAnalyzableServersCmd).Standalone()

		migrationhubstrategy_listAnalyzableServersCmd.Flags().String("max-results", "", "The maximum number of items to include in the response.")
		migrationhubstrategy_listAnalyzableServersCmd.Flags().String("next-token", "", "The token from a previous call that you use to retrieve the next set of results.")
		migrationhubstrategy_listAnalyzableServersCmd.Flags().String("sort", "", "Specifies whether to sort by ascending (ASC) or descending (DESC) order.")
	})
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_listAnalyzableServersCmd)
}
