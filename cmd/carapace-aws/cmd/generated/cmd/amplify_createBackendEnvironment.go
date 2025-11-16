package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_createBackendEnvironmentCmd = &cobra.Command{
	Use:   "create-backend-environment",
	Short: "Creates a new backend environment for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_createBackendEnvironmentCmd).Standalone()

	amplify_createBackendEnvironmentCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
	amplify_createBackendEnvironmentCmd.Flags().String("deployment-artifacts", "", "The name of deployment artifacts.")
	amplify_createBackendEnvironmentCmd.Flags().String("environment-name", "", "The name for the backend environment.")
	amplify_createBackendEnvironmentCmd.Flags().String("stack-name", "", "The AWS CloudFormation stack name of a backend environment.")
	amplify_createBackendEnvironmentCmd.MarkFlagRequired("app-id")
	amplify_createBackendEnvironmentCmd.MarkFlagRequired("environment-name")
	amplifyCmd.AddCommand(amplify_createBackendEnvironmentCmd)
}
