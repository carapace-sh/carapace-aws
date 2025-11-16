package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_deleteThemeCmd = &cobra.Command{
	Use:   "delete-theme",
	Short: "Deletes a theme from an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_deleteThemeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifyuibuilder_deleteThemeCmd).Standalone()

		amplifyuibuilder_deleteThemeCmd.Flags().String("app-id", "", "The unique ID of the Amplify app associated with the theme to delete.")
		amplifyuibuilder_deleteThemeCmd.Flags().String("environment-name", "", "The name of the backend environment that is a part of the Amplify app.")
		amplifyuibuilder_deleteThemeCmd.Flags().String("id", "", "The unique ID of the theme to delete.")
		amplifyuibuilder_deleteThemeCmd.MarkFlagRequired("app-id")
		amplifyuibuilder_deleteThemeCmd.MarkFlagRequired("environment-name")
		amplifyuibuilder_deleteThemeCmd.MarkFlagRequired("id")
	})
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_deleteThemeCmd)
}
