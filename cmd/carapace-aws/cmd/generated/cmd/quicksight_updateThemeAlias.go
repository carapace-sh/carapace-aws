package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateThemeAliasCmd = &cobra.Command{
	Use:   "update-theme-alias",
	Short: "Updates an alias of a theme.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateThemeAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateThemeAliasCmd).Standalone()

		quicksight_updateThemeAliasCmd.Flags().String("alias-name", "", "The name of the theme alias that you want to update.")
		quicksight_updateThemeAliasCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the theme alias that you're updating.")
		quicksight_updateThemeAliasCmd.Flags().String("theme-id", "", "The ID for the theme.")
		quicksight_updateThemeAliasCmd.Flags().String("theme-version-number", "", "The version number of the theme that the alias should reference.")
		quicksight_updateThemeAliasCmd.MarkFlagRequired("alias-name")
		quicksight_updateThemeAliasCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateThemeAliasCmd.MarkFlagRequired("theme-id")
		quicksight_updateThemeAliasCmd.MarkFlagRequired("theme-version-number")
	})
	quicksightCmd.AddCommand(quicksight_updateThemeAliasCmd)
}
