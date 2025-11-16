package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_listRevisionAssetsCmd = &cobra.Command{
	Use:   "list-revision-assets",
	Short: "This operation lists a revision's assets sorted alphabetically in descending order.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_listRevisionAssetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_listRevisionAssetsCmd).Standalone()

		dataexchange_listRevisionAssetsCmd.Flags().String("data-set-id", "", "The unique identifier for a data set.")
		dataexchange_listRevisionAssetsCmd.Flags().String("max-results", "", "The maximum number of results returned by a single call.")
		dataexchange_listRevisionAssetsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
		dataexchange_listRevisionAssetsCmd.Flags().String("revision-id", "", "The unique identifier for a revision.")
		dataexchange_listRevisionAssetsCmd.MarkFlagRequired("data-set-id")
		dataexchange_listRevisionAssetsCmd.MarkFlagRequired("revision-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_listRevisionAssetsCmd)
}
