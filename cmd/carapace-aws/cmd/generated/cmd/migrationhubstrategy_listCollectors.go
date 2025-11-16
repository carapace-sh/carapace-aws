package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_listCollectorsCmd = &cobra.Command{
	Use:   "list-collectors",
	Short: "Retrieves a list of all the installed collectors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_listCollectorsCmd).Standalone()

	migrationhubstrategy_listCollectorsCmd.Flags().String("max-results", "", "The maximum number of items to include in the response.")
	migrationhubstrategy_listCollectorsCmd.Flags().String("next-token", "", "The token from a previous call that you use to retrieve the next set of results.")
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_listCollectorsCmd)
}
