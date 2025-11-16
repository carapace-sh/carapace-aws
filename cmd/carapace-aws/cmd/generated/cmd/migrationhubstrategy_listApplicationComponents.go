package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_listApplicationComponentsCmd = &cobra.Command{
	Use:   "list-application-components",
	Short: "Retrieves a list of all the application components (processes).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_listApplicationComponentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubstrategy_listApplicationComponentsCmd).Standalone()

		migrationhubstrategy_listApplicationComponentsCmd.Flags().String("application-component-criteria", "", "Criteria for filtering the list of application components.")
		migrationhubstrategy_listApplicationComponentsCmd.Flags().String("filter-value", "", "Specify the value based on the application component criteria type.")
		migrationhubstrategy_listApplicationComponentsCmd.Flags().String("group-id-filter", "", "The group ID specified in to filter on.")
		migrationhubstrategy_listApplicationComponentsCmd.Flags().String("max-results", "", "The maximum number of items to include in the response.")
		migrationhubstrategy_listApplicationComponentsCmd.Flags().String("next-token", "", "The token from a previous call that you use to retrieve the next set of results.")
		migrationhubstrategy_listApplicationComponentsCmd.Flags().String("sort", "", "Specifies whether to sort by ascending (`ASC`) or descending (`DESC`) order.")
	})
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_listApplicationComponentsCmd)
}
