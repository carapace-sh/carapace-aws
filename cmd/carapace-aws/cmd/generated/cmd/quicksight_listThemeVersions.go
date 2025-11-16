package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listThemeVersionsCmd = &cobra.Command{
	Use:   "list-theme-versions",
	Short: "Lists all the versions of the themes in the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listThemeVersionsCmd).Standalone()

	quicksight_listThemeVersionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the themes that you're listing.")
	quicksight_listThemeVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_listThemeVersionsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
	quicksight_listThemeVersionsCmd.Flags().String("theme-id", "", "The ID for the theme.")
	quicksight_listThemeVersionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_listThemeVersionsCmd.MarkFlagRequired("theme-id")
	quicksightCmd.AddCommand(quicksight_listThemeVersionsCmd)
}
