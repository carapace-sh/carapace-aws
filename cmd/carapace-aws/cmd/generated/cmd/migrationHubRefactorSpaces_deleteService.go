package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_deleteServiceCmd = &cobra.Command{
	Use:   "delete-service",
	Short: "Deletes an Amazon Web Services Migration Hub Refactor Spaces service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_deleteServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpaces_deleteServiceCmd).Standalone()

		migrationHubRefactorSpaces_deleteServiceCmd.Flags().String("application-identifier", "", "Deletes a Refactor Spaces service.")
		migrationHubRefactorSpaces_deleteServiceCmd.Flags().String("environment-identifier", "", "The ID of the environment that the service is in.")
		migrationHubRefactorSpaces_deleteServiceCmd.Flags().String("service-identifier", "", "The ID of the service to delete.")
		migrationHubRefactorSpaces_deleteServiceCmd.MarkFlagRequired("application-identifier")
		migrationHubRefactorSpaces_deleteServiceCmd.MarkFlagRequired("environment-identifier")
		migrationHubRefactorSpaces_deleteServiceCmd.MarkFlagRequired("service-identifier")
	})
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_deleteServiceCmd)
}
