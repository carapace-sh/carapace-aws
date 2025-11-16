package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_listComponentsCmd = &cobra.Command{
	Use:   "list-components",
	Short: "Retrieves a list of components for a specified Amplify app and backend environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_listComponentsCmd).Standalone()

	amplifyuibuilder_listComponentsCmd.Flags().String("app-id", "", "The unique ID for the Amplify app.")
	amplifyuibuilder_listComponentsCmd.Flags().String("environment-name", "", "The name of the backend environment that is a part of the Amplify app.")
	amplifyuibuilder_listComponentsCmd.Flags().String("max-results", "", "The maximum number of components to retrieve.")
	amplifyuibuilder_listComponentsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	amplifyuibuilder_listComponentsCmd.MarkFlagRequired("app-id")
	amplifyuibuilder_listComponentsCmd.MarkFlagRequired("environment-name")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_listComponentsCmd)
}
