package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_updateThemeCmd = &cobra.Command{
	Use:   "update-theme",
	Short: "Updates an existing theme.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_updateThemeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifyuibuilder_updateThemeCmd).Standalone()

		amplifyuibuilder_updateThemeCmd.Flags().String("app-id", "", "The unique ID for the Amplify app.")
		amplifyuibuilder_updateThemeCmd.Flags().String("client-token", "", "The unique client token.")
		amplifyuibuilder_updateThemeCmd.Flags().String("environment-name", "", "The name of the backend environment that is part of the Amplify app.")
		amplifyuibuilder_updateThemeCmd.Flags().String("id", "", "The unique ID for the theme.")
		amplifyuibuilder_updateThemeCmd.Flags().String("updated-theme", "", "The configuration of the updated theme.")
		amplifyuibuilder_updateThemeCmd.MarkFlagRequired("app-id")
		amplifyuibuilder_updateThemeCmd.MarkFlagRequired("environment-name")
		amplifyuibuilder_updateThemeCmd.MarkFlagRequired("id")
		amplifyuibuilder_updateThemeCmd.MarkFlagRequired("updated-theme")
	})
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_updateThemeCmd)
}
