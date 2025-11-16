package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_getServiceCmd = &cobra.Command{
	Use:   "get-service",
	Short: "Gets an Amazon Web Services Migration Hub Refactor Spaces service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_getServiceCmd).Standalone()

	migrationHubRefactorSpaces_getServiceCmd.Flags().String("application-identifier", "", "The ID of the application.")
	migrationHubRefactorSpaces_getServiceCmd.Flags().String("environment-identifier", "", "The ID of the environment.")
	migrationHubRefactorSpaces_getServiceCmd.Flags().String("service-identifier", "", "The ID of the service.")
	migrationHubRefactorSpaces_getServiceCmd.MarkFlagRequired("application-identifier")
	migrationHubRefactorSpaces_getServiceCmd.MarkFlagRequired("environment-identifier")
	migrationHubRefactorSpaces_getServiceCmd.MarkFlagRequired("service-identifier")
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_getServiceCmd)
}
