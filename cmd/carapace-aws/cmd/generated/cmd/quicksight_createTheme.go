package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createThemeCmd = &cobra.Command{
	Use:   "create-theme",
	Short: "Creates a theme.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createThemeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createThemeCmd).Standalone()

		quicksight_createThemeCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account where you want to store the new theme.")
		quicksight_createThemeCmd.Flags().String("base-theme-id", "", "The ID of the theme that a custom theme will inherit from.")
		quicksight_createThemeCmd.Flags().String("configuration", "", "The theme configuration, which contains the theme display properties.")
		quicksight_createThemeCmd.Flags().String("name", "", "A display name for the theme.")
		quicksight_createThemeCmd.Flags().String("permissions", "", "A valid grouping of resource permissions to apply to the new theme.")
		quicksight_createThemeCmd.Flags().String("tags", "", "A map of the key-value pairs for the resource tag or tags that you want to add to the resource.")
		quicksight_createThemeCmd.Flags().String("theme-id", "", "An ID for the theme that you want to create.")
		quicksight_createThemeCmd.Flags().String("version-description", "", "A description of the first version of the theme that you're creating.")
		quicksight_createThemeCmd.MarkFlagRequired("aws-account-id")
		quicksight_createThemeCmd.MarkFlagRequired("base-theme-id")
		quicksight_createThemeCmd.MarkFlagRequired("configuration")
		quicksight_createThemeCmd.MarkFlagRequired("name")
		quicksight_createThemeCmd.MarkFlagRequired("theme-id")
	})
	quicksightCmd.AddCommand(quicksight_createThemeCmd)
}
