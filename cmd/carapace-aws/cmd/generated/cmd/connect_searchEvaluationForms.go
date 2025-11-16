package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchEvaluationFormsCmd = &cobra.Command{
	Use:   "search-evaluation-forms",
	Short: "Searches evaluation forms in an Amazon Connect instance, with optional filtering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchEvaluationFormsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_searchEvaluationFormsCmd).Standalone()

		connect_searchEvaluationFormsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_searchEvaluationFormsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_searchEvaluationFormsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_searchEvaluationFormsCmd.Flags().String("search-criteria", "", "The search criteria to be used to return evaluation forms.")
		connect_searchEvaluationFormsCmd.Flags().String("search-filter", "", "Filters to be applied to search results.")
		connect_searchEvaluationFormsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_searchEvaluationFormsCmd)
}
