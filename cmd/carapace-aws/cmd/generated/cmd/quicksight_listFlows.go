package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listFlowsCmd = &cobra.Command{
	Use:   "list-flows",
	Short: "Lists flows in an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listFlowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listFlowsCmd).Standalone()

		quicksight_listFlowsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the flow list that you are getting.")
		quicksight_listFlowsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		quicksight_listFlowsCmd.Flags().String("next-token", "", "The token to request the next set of results, or null if you want to retrieve the first set.")
		quicksight_listFlowsCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_listFlowsCmd)
}
