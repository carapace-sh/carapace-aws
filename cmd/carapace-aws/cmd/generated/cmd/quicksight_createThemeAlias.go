package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createThemeAliasCmd = &cobra.Command{
	Use:   "create-theme-alias",
	Short: "Creates a theme alias for a theme.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createThemeAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createThemeAliasCmd).Standalone()

		quicksight_createThemeAliasCmd.Flags().String("alias-name", "", "The name that you want to give to the theme alias that you are creating.")
		quicksight_createThemeAliasCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the theme for the new theme alias.")
		quicksight_createThemeAliasCmd.Flags().String("theme-id", "", "An ID for the theme alias.")
		quicksight_createThemeAliasCmd.Flags().String("theme-version-number", "", "The version number of the theme.")
		quicksight_createThemeAliasCmd.MarkFlagRequired("alias-name")
		quicksight_createThemeAliasCmd.MarkFlagRequired("aws-account-id")
		quicksight_createThemeAliasCmd.MarkFlagRequired("theme-id")
		quicksight_createThemeAliasCmd.MarkFlagRequired("theme-version-number")
	})
	quicksightCmd.AddCommand(quicksight_createThemeAliasCmd)
}
