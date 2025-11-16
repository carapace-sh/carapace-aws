package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateThemeCmd = &cobra.Command{
	Use:   "update-theme",
	Short: "Updates a theme.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateThemeCmd).Standalone()

	quicksight_updateThemeCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the theme that you're updating.")
	quicksight_updateThemeCmd.Flags().String("base-theme-id", "", "The theme ID, defined by Amazon Quick Sight, that a custom theme inherits from.")
	quicksight_updateThemeCmd.Flags().String("configuration", "", "The theme configuration, which contains the theme display properties.")
	quicksight_updateThemeCmd.Flags().String("name", "", "The name for the theme.")
	quicksight_updateThemeCmd.Flags().String("theme-id", "", "The ID for the theme.")
	quicksight_updateThemeCmd.Flags().String("version-description", "", "A description of the theme version that you're updating Every time that you call `UpdateTheme`, you create a new version of the theme.")
	quicksight_updateThemeCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateThemeCmd.MarkFlagRequired("base-theme-id")
	quicksight_updateThemeCmd.MarkFlagRequired("theme-id")
	quicksightCmd.AddCommand(quicksight_updateThemeCmd)
}
