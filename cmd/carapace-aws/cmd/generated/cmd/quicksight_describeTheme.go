package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeThemeCmd = &cobra.Command{
	Use:   "describe-theme",
	Short: "Describes a theme.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeThemeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeThemeCmd).Standalone()

		quicksight_describeThemeCmd.Flags().String("alias-name", "", "The alias of the theme that you want to describe.")
		quicksight_describeThemeCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the theme that you're describing.")
		quicksight_describeThemeCmd.Flags().String("theme-id", "", "The ID for the theme.")
		quicksight_describeThemeCmd.Flags().String("version-number", "", "The version number for the version to describe.")
		quicksight_describeThemeCmd.MarkFlagRequired("aws-account-id")
		quicksight_describeThemeCmd.MarkFlagRequired("theme-id")
	})
	quicksightCmd.AddCommand(quicksight_describeThemeCmd)
}
