package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_listDataSetRevisionsCmd = &cobra.Command{
	Use:   "list-data-set-revisions",
	Short: "This operation lists a data set's revisions sorted by CreatedAt in descending order.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_listDataSetRevisionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_listDataSetRevisionsCmd).Standalone()

		dataexchange_listDataSetRevisionsCmd.Flags().String("data-set-id", "", "The unique identifier for a data set.")
		dataexchange_listDataSetRevisionsCmd.Flags().String("max-results", "", "The maximum number of results returned by a single call.")
		dataexchange_listDataSetRevisionsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
		dataexchange_listDataSetRevisionsCmd.MarkFlagRequired("data-set-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_listDataSetRevisionsCmd)
}
