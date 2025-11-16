package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchContactEvaluationsCmd = &cobra.Command{
	Use:   "search-contact-evaluations",
	Short: "Searches contact evaluations in an Amazon Connect instance, with optional filtering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchContactEvaluationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_searchContactEvaluationsCmd).Standalone()

		connect_searchContactEvaluationsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_searchContactEvaluationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_searchContactEvaluationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_searchContactEvaluationsCmd.Flags().String("search-criteria", "", "The search criteria to be used to return contact evaluations.")
		connect_searchContactEvaluationsCmd.Flags().String("search-filter", "", "Filters to be applied to search results.")
		connect_searchContactEvaluationsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_searchContactEvaluationsCmd)
}
