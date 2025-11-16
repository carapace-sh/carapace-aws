package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_listFormsCmd = &cobra.Command{
	Use:   "list-forms",
	Short: "Retrieves a list of forms for a specified Amplify app and backend environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_listFormsCmd).Standalone()

	amplifyuibuilder_listFormsCmd.Flags().String("app-id", "", "The unique ID for the Amplify app.")
	amplifyuibuilder_listFormsCmd.Flags().String("environment-name", "", "The name of the backend environment that is a part of the Amplify app.")
	amplifyuibuilder_listFormsCmd.Flags().String("max-results", "", "The maximum number of forms to retrieve.")
	amplifyuibuilder_listFormsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	amplifyuibuilder_listFormsCmd.MarkFlagRequired("app-id")
	amplifyuibuilder_listFormsCmd.MarkFlagRequired("environment-name")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_listFormsCmd)
}
