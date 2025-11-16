package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_listThemesCmd = &cobra.Command{
	Use:   "list-themes",
	Short: "Retrieves a list of themes for a specified Amplify app and backend environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_listThemesCmd).Standalone()

	amplifyuibuilder_listThemesCmd.Flags().String("app-id", "", "The unique ID for the Amplify app.")
	amplifyuibuilder_listThemesCmd.Flags().String("environment-name", "", "The name of the backend environment that is a part of the Amplify app.")
	amplifyuibuilder_listThemesCmd.Flags().String("max-results", "", "The maximum number of theme results to return in the response.")
	amplifyuibuilder_listThemesCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	amplifyuibuilder_listThemesCmd.MarkFlagRequired("app-id")
	amplifyuibuilder_listThemesCmd.MarkFlagRequired("environment-name")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_listThemesCmd)
}
