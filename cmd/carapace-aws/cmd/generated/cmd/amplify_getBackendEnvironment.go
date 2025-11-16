package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_getBackendEnvironmentCmd = &cobra.Command{
	Use:   "get-backend-environment",
	Short: "Returns a backend environment for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_getBackendEnvironmentCmd).Standalone()

	amplify_getBackendEnvironmentCmd.Flags().String("app-id", "", "The unique id for an Amplify app.")
	amplify_getBackendEnvironmentCmd.Flags().String("environment-name", "", "The name for the backend environment.")
	amplify_getBackendEnvironmentCmd.MarkFlagRequired("app-id")
	amplify_getBackendEnvironmentCmd.MarkFlagRequired("environment-name")
	amplifyCmd.AddCommand(amplify_getBackendEnvironmentCmd)
}
