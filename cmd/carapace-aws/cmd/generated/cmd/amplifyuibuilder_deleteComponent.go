package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_deleteComponentCmd = &cobra.Command{
	Use:   "delete-component",
	Short: "Deletes a component from an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_deleteComponentCmd).Standalone()

	amplifyuibuilder_deleteComponentCmd.Flags().String("app-id", "", "The unique ID of the Amplify app associated with the component to delete.")
	amplifyuibuilder_deleteComponentCmd.Flags().String("environment-name", "", "The name of the backend environment that is a part of the Amplify app.")
	amplifyuibuilder_deleteComponentCmd.Flags().String("id", "", "The unique ID of the component to delete.")
	amplifyuibuilder_deleteComponentCmd.MarkFlagRequired("app-id")
	amplifyuibuilder_deleteComponentCmd.MarkFlagRequired("environment-name")
	amplifyuibuilder_deleteComponentCmd.MarkFlagRequired("id")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_deleteComponentCmd)
}
