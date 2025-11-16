package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "This operation lists your jobs sorted by CreatedAt in descending order.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_listJobsCmd).Standalone()

	dataexchange_listJobsCmd.Flags().String("data-set-id", "", "The unique identifier for a data set.")
	dataexchange_listJobsCmd.Flags().String("max-results", "", "The maximum number of results returned by a single call.")
	dataexchange_listJobsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
	dataexchange_listJobsCmd.Flags().String("revision-id", "", "The unique identifier for a revision.")
	dataexchangeCmd.AddCommand(dataexchange_listJobsCmd)
}
