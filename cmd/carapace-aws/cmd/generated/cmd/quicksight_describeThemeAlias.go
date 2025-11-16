package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeThemeAliasCmd = &cobra.Command{
	Use:   "describe-theme-alias",
	Short: "Describes the alias for a theme.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeThemeAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeThemeAliasCmd).Standalone()

		quicksight_describeThemeAliasCmd.Flags().String("alias-name", "", "The name of the theme alias that you want to describe.")
		quicksight_describeThemeAliasCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the theme alias that you're describing.")
		quicksight_describeThemeAliasCmd.Flags().String("theme-id", "", "The ID for the theme.")
		quicksight_describeThemeAliasCmd.MarkFlagRequired("alias-name")
		quicksight_describeThemeAliasCmd.MarkFlagRequired("aws-account-id")
		quicksight_describeThemeAliasCmd.MarkFlagRequired("theme-id")
	})
	quicksightCmd.AddCommand(quicksight_describeThemeAliasCmd)
}
