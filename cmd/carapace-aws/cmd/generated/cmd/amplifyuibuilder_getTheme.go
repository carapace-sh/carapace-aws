package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_getThemeCmd = &cobra.Command{
	Use:   "get-theme",
	Short: "Returns an existing theme for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_getThemeCmd).Standalone()

	amplifyuibuilder_getThemeCmd.Flags().String("app-id", "", "The unique ID of the Amplify app.")
	amplifyuibuilder_getThemeCmd.Flags().String("environment-name", "", "The name of the backend environment that is part of the Amplify app.")
	amplifyuibuilder_getThemeCmd.Flags().String("id", "", "The unique ID for the theme.")
	amplifyuibuilder_getThemeCmd.MarkFlagRequired("app-id")
	amplifyuibuilder_getThemeCmd.MarkFlagRequired("environment-name")
	amplifyuibuilder_getThemeCmd.MarkFlagRequired("id")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_getThemeCmd)
}
