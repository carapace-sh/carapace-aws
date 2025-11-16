package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteThemeCmd = &cobra.Command{
	Use:   "delete-theme",
	Short: "Deletes a theme.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteThemeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteThemeCmd).Standalone()

		quicksight_deleteThemeCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the theme that you're deleting.")
		quicksight_deleteThemeCmd.Flags().String("theme-id", "", "An ID for the theme that you want to delete.")
		quicksight_deleteThemeCmd.Flags().String("version-number", "", "The version of the theme that you want to delete.")
		quicksight_deleteThemeCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteThemeCmd.MarkFlagRequired("theme-id")
	})
	quicksightCmd.AddCommand(quicksight_deleteThemeCmd)
}
