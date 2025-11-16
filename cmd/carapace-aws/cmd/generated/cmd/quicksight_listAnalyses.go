package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listAnalysesCmd = &cobra.Command{
	Use:   "list-analyses",
	Short: "Lists Amazon Quick Sight analyses that exist in the specified Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listAnalysesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listAnalysesCmd).Standalone()

		quicksight_listAnalysesCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the analyses.")
		quicksight_listAnalysesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		quicksight_listAnalysesCmd.Flags().String("next-token", "", "A pagination token that can be used in a subsequent request.")
		quicksight_listAnalysesCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_listAnalysesCmd)
}
