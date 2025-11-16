package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listThemesCmd = &cobra.Command{
	Use:   "list-themes",
	Short: "Lists all the themes in the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listThemesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listThemesCmd).Standalone()

		quicksight_listThemesCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the themes that you're listing.")
		quicksight_listThemesCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		quicksight_listThemesCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
		quicksight_listThemesCmd.Flags().String("type", "", "The type of themes that you want to list.")
		quicksight_listThemesCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_listThemesCmd)
}
