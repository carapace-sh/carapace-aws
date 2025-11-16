package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_getFormCmd = &cobra.Command{
	Use:   "get-form",
	Short: "Returns an existing form for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_getFormCmd).Standalone()

	amplifyuibuilder_getFormCmd.Flags().String("app-id", "", "The unique ID of the Amplify app.")
	amplifyuibuilder_getFormCmd.Flags().String("environment-name", "", "The name of the backend environment that is part of the Amplify app.")
	amplifyuibuilder_getFormCmd.Flags().String("id", "", "The unique ID of the form.")
	amplifyuibuilder_getFormCmd.MarkFlagRequired("app-id")
	amplifyuibuilder_getFormCmd.MarkFlagRequired("environment-name")
	amplifyuibuilder_getFormCmd.MarkFlagRequired("id")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_getFormCmd)
}
