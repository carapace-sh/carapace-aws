package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_listServersCmd = &cobra.Command{
	Use:   "list-servers",
	Short: "Returns a list of all the servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_listServersCmd).Standalone()

	migrationhubstrategy_listServersCmd.Flags().String("filter-value", "", "Specifies the filter value, which is based on the type of server criteria.")
	migrationhubstrategy_listServersCmd.Flags().String("group-id-filter", "", "Specifies the group ID to filter on.")
	migrationhubstrategy_listServersCmd.Flags().String("max-results", "", "The maximum number of items to include in the response.")
	migrationhubstrategy_listServersCmd.Flags().String("next-token", "", "The token from a previous call that you use to retrieve the next set of results.")
	migrationhubstrategy_listServersCmd.Flags().String("server-criteria", "", "Criteria for filtering servers.")
	migrationhubstrategy_listServersCmd.Flags().String("sort", "", "Specifies whether to sort by ascending (`ASC`) or descending (`DESC`) order.")
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_listServersCmd)
}
