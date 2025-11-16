package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listBrandsCmd = &cobra.Command{
	Use:   "list-brands",
	Short: "Lists all brands in an Quick Sight account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listBrandsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listBrandsCmd).Standalone()

		quicksight_listBrandsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that owns the brands that you want to list.")
		quicksight_listBrandsCmd.Flags().String("max-results", "", "The maximum number of results to be returned in a single request.")
		quicksight_listBrandsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
		quicksight_listBrandsCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_listBrandsCmd)
}
