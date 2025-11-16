package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_createFormCmd = &cobra.Command{
	Use:   "create-form",
	Short: "Creates a new form for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_createFormCmd).Standalone()

	amplifyuibuilder_createFormCmd.Flags().String("app-id", "", "The unique ID of the Amplify app to associate with the form.")
	amplifyuibuilder_createFormCmd.Flags().String("client-token", "", "The unique client token.")
	amplifyuibuilder_createFormCmd.Flags().String("environment-name", "", "The name of the backend environment that is a part of the Amplify app.")
	amplifyuibuilder_createFormCmd.Flags().String("form-to-create", "", "Represents the configuration of the form to create.")
	amplifyuibuilder_createFormCmd.MarkFlagRequired("app-id")
	amplifyuibuilder_createFormCmd.MarkFlagRequired("environment-name")
	amplifyuibuilder_createFormCmd.MarkFlagRequired("form-to-create")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_createFormCmd)
}
