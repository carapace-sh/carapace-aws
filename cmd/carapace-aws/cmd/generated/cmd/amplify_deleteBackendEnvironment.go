package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_deleteBackendEnvironmentCmd = &cobra.Command{
	Use:   "delete-backend-environment",
	Short: "Deletes a backend environment for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_deleteBackendEnvironmentCmd).Standalone()

	amplify_deleteBackendEnvironmentCmd.Flags().String("app-id", "", "The unique ID of an Amplify app.")
	amplify_deleteBackendEnvironmentCmd.Flags().String("environment-name", "", "The name of a backend environment of an Amplify app.")
	amplify_deleteBackendEnvironmentCmd.MarkFlagRequired("app-id")
	amplify_deleteBackendEnvironmentCmd.MarkFlagRequired("environment-name")
	amplifyCmd.AddCommand(amplify_deleteBackendEnvironmentCmd)
}
