package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listThemeAliasesCmd = &cobra.Command{
	Use:   "list-theme-aliases",
	Short: "Lists all the aliases of a theme.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listThemeAliasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listThemeAliasesCmd).Standalone()

		quicksight_listThemeAliasesCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the theme aliases that you're listing.")
		quicksight_listThemeAliasesCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		quicksight_listThemeAliasesCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
		quicksight_listThemeAliasesCmd.Flags().String("theme-id", "", "The ID for the theme.")
		quicksight_listThemeAliasesCmd.MarkFlagRequired("aws-account-id")
		quicksight_listThemeAliasesCmd.MarkFlagRequired("theme-id")
	})
	quicksightCmd.AddCommand(quicksight_listThemeAliasesCmd)
}
