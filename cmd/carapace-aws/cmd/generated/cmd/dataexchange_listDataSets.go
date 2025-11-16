package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_listDataSetsCmd = &cobra.Command{
	Use:   "list-data-sets",
	Short: "This operation lists your data sets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_listDataSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_listDataSetsCmd).Standalone()

		dataexchange_listDataSetsCmd.Flags().String("max-results", "", "The maximum number of results returned by a single call.")
		dataexchange_listDataSetsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
		dataexchange_listDataSetsCmd.Flags().String("origin", "", "A property that defines the data set as OWNED by the account (for providers) or ENTITLED to the account (for subscribers).")
	})
	dataexchangeCmd.AddCommand(dataexchange_listDataSetsCmd)
}
