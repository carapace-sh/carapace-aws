package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_updateComponentCmd = &cobra.Command{
	Use:   "update-component",
	Short: "Updates an existing component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_updateComponentCmd).Standalone()

	amplifyuibuilder_updateComponentCmd.Flags().String("app-id", "", "The unique ID for the Amplify app.")
	amplifyuibuilder_updateComponentCmd.Flags().String("client-token", "", "The unique client token.")
	amplifyuibuilder_updateComponentCmd.Flags().String("environment-name", "", "The name of the backend environment that is part of the Amplify app.")
	amplifyuibuilder_updateComponentCmd.Flags().String("id", "", "The unique ID for the component.")
	amplifyuibuilder_updateComponentCmd.Flags().String("updated-component", "", "The configuration of the updated component.")
	amplifyuibuilder_updateComponentCmd.MarkFlagRequired("app-id")
	amplifyuibuilder_updateComponentCmd.MarkFlagRequired("environment-name")
	amplifyuibuilder_updateComponentCmd.MarkFlagRequired("id")
	amplifyuibuilder_updateComponentCmd.MarkFlagRequired("updated-component")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_updateComponentCmd)
}
