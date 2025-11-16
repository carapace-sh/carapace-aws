package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_createEnvironmentCmd = &cobra.Command{
	Use:   "create-environment",
	Short: "Creates an Amazon Web Services Migration Hub Refactor Spaces environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_createEnvironmentCmd).Standalone()

	migrationHubRefactorSpaces_createEnvironmentCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	migrationHubRefactorSpaces_createEnvironmentCmd.Flags().String("description", "", "The description of the environment.")
	migrationHubRefactorSpaces_createEnvironmentCmd.Flags().String("name", "", "The name of the environment.")
	migrationHubRefactorSpaces_createEnvironmentCmd.Flags().String("network-fabric-type", "", "The network fabric type of the environment.")
	migrationHubRefactorSpaces_createEnvironmentCmd.Flags().String("tags", "", "The tags to assign to the environment.")
	migrationHubRefactorSpaces_createEnvironmentCmd.MarkFlagRequired("name")
	migrationHubRefactorSpaces_createEnvironmentCmd.MarkFlagRequired("network-fabric-type")
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_createEnvironmentCmd)
}
