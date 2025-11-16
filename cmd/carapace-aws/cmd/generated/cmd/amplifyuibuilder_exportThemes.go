package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_exportThemesCmd = &cobra.Command{
	Use:   "export-themes",
	Short: "Exports theme configurations to code that is ready to integrate into an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_exportThemesCmd).Standalone()

	amplifyuibuilder_exportThemesCmd.Flags().String("app-id", "", "The unique ID of the Amplify app to export the themes to.")
	amplifyuibuilder_exportThemesCmd.Flags().String("environment-name", "", "The name of the backend environment that is part of the Amplify app.")
	amplifyuibuilder_exportThemesCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	amplifyuibuilder_exportThemesCmd.MarkFlagRequired("app-id")
	amplifyuibuilder_exportThemesCmd.MarkFlagRequired("environment-name")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_exportThemesCmd)
}
