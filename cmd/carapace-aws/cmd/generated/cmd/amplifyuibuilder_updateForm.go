package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_updateFormCmd = &cobra.Command{
	Use:   "update-form",
	Short: "Updates an existing form.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_updateFormCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifyuibuilder_updateFormCmd).Standalone()

		amplifyuibuilder_updateFormCmd.Flags().String("app-id", "", "The unique ID for the Amplify app.")
		amplifyuibuilder_updateFormCmd.Flags().String("client-token", "", "The unique client token.")
		amplifyuibuilder_updateFormCmd.Flags().String("environment-name", "", "The name of the backend environment that is part of the Amplify app.")
		amplifyuibuilder_updateFormCmd.Flags().String("id", "", "The unique ID for the form.")
		amplifyuibuilder_updateFormCmd.Flags().String("updated-form", "", "The request accepts the following data in JSON format.")
		amplifyuibuilder_updateFormCmd.MarkFlagRequired("app-id")
		amplifyuibuilder_updateFormCmd.MarkFlagRequired("environment-name")
		amplifyuibuilder_updateFormCmd.MarkFlagRequired("id")
		amplifyuibuilder_updateFormCmd.MarkFlagRequired("updated-form")
	})
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_updateFormCmd)
}
