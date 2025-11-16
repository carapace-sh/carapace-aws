package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_exportComponentsCmd = &cobra.Command{
	Use:   "export-components",
	Short: "Exports component configurations to code that is ready to integrate into an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_exportComponentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifyuibuilder_exportComponentsCmd).Standalone()

		amplifyuibuilder_exportComponentsCmd.Flags().String("app-id", "", "The unique ID of the Amplify app to export components to.")
		amplifyuibuilder_exportComponentsCmd.Flags().String("environment-name", "", "The name of the backend environment that is a part of the Amplify app.")
		amplifyuibuilder_exportComponentsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
		amplifyuibuilder_exportComponentsCmd.MarkFlagRequired("app-id")
		amplifyuibuilder_exportComponentsCmd.MarkFlagRequired("environment-name")
	})
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_exportComponentsCmd)
}
