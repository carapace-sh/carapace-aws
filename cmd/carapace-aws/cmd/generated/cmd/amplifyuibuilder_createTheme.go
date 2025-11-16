package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_createThemeCmd = &cobra.Command{
	Use:   "create-theme",
	Short: "Creates a theme to apply to the components in an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_createThemeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifyuibuilder_createThemeCmd).Standalone()

		amplifyuibuilder_createThemeCmd.Flags().String("app-id", "", "The unique ID of the Amplify app associated with the theme.")
		amplifyuibuilder_createThemeCmd.Flags().String("client-token", "", "The unique client token.")
		amplifyuibuilder_createThemeCmd.Flags().String("environment-name", "", "The name of the backend environment that is a part of the Amplify app.")
		amplifyuibuilder_createThemeCmd.Flags().String("theme-to-create", "", "Represents the configuration of the theme to create.")
		amplifyuibuilder_createThemeCmd.MarkFlagRequired("app-id")
		amplifyuibuilder_createThemeCmd.MarkFlagRequired("environment-name")
		amplifyuibuilder_createThemeCmd.MarkFlagRequired("theme-to-create")
	})
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_createThemeCmd)
}
