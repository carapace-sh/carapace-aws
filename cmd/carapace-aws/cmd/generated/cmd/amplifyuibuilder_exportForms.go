package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_exportFormsCmd = &cobra.Command{
	Use:   "export-forms",
	Short: "Exports form configurations to code that is ready to integrate into an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_exportFormsCmd).Standalone()

	amplifyuibuilder_exportFormsCmd.Flags().String("app-id", "", "The unique ID of the Amplify app to export forms to.")
	amplifyuibuilder_exportFormsCmd.Flags().String("environment-name", "", "The name of the backend environment that is a part of the Amplify app.")
	amplifyuibuilder_exportFormsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	amplifyuibuilder_exportFormsCmd.MarkFlagRequired("app-id")
	amplifyuibuilder_exportFormsCmd.MarkFlagRequired("environment-name")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_exportFormsCmd)
}
