package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_deleteFormCmd = &cobra.Command{
	Use:   "delete-form",
	Short: "Deletes a form from an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_deleteFormCmd).Standalone()

	amplifyuibuilder_deleteFormCmd.Flags().String("app-id", "", "The unique ID of the Amplify app associated with the form to delete.")
	amplifyuibuilder_deleteFormCmd.Flags().String("environment-name", "", "The name of the backend environment that is a part of the Amplify app.")
	amplifyuibuilder_deleteFormCmd.Flags().String("id", "", "The unique ID of the form to delete.")
	amplifyuibuilder_deleteFormCmd.MarkFlagRequired("app-id")
	amplifyuibuilder_deleteFormCmd.MarkFlagRequired("environment-name")
	amplifyuibuilder_deleteFormCmd.MarkFlagRequired("id")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_deleteFormCmd)
}
