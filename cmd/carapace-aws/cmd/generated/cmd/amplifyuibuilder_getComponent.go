package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_getComponentCmd = &cobra.Command{
	Use:   "get-component",
	Short: "Returns an existing component for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_getComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifyuibuilder_getComponentCmd).Standalone()

		amplifyuibuilder_getComponentCmd.Flags().String("app-id", "", "The unique ID of the Amplify app.")
		amplifyuibuilder_getComponentCmd.Flags().String("environment-name", "", "The name of the backend environment that is part of the Amplify app.")
		amplifyuibuilder_getComponentCmd.Flags().String("id", "", "The unique ID of the component.")
		amplifyuibuilder_getComponentCmd.MarkFlagRequired("app-id")
		amplifyuibuilder_getComponentCmd.MarkFlagRequired("environment-name")
		amplifyuibuilder_getComponentCmd.MarkFlagRequired("id")
	})
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_getComponentCmd)
}
