package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_createComponentCmd = &cobra.Command{
	Use:   "create-component",
	Short: "Creates a new component for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_createComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifyuibuilder_createComponentCmd).Standalone()

		amplifyuibuilder_createComponentCmd.Flags().String("app-id", "", "The unique ID of the Amplify app to associate with the component.")
		amplifyuibuilder_createComponentCmd.Flags().String("client-token", "", "The unique client token.")
		amplifyuibuilder_createComponentCmd.Flags().String("component-to-create", "", "Represents the configuration of the component to create.")
		amplifyuibuilder_createComponentCmd.Flags().String("environment-name", "", "The name of the backend environment that is a part of the Amplify app.")
		amplifyuibuilder_createComponentCmd.MarkFlagRequired("app-id")
		amplifyuibuilder_createComponentCmd.MarkFlagRequired("component-to-create")
		amplifyuibuilder_createComponentCmd.MarkFlagRequired("environment-name")
	})
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_createComponentCmd)
}
