package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_createServiceCmd = &cobra.Command{
	Use:   "create-service",
	Short: "Creates an Amazon Web Services Migration Hub Refactor Spaces service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_createServiceCmd).Standalone()

	migrationHubRefactorSpaces_createServiceCmd.Flags().String("application-identifier", "", "The ID of the application which the service is created.")
	migrationHubRefactorSpaces_createServiceCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	migrationHubRefactorSpaces_createServiceCmd.Flags().String("description", "", "The description of the service.")
	migrationHubRefactorSpaces_createServiceCmd.Flags().String("endpoint-type", "", "The type of endpoint to use for the service.")
	migrationHubRefactorSpaces_createServiceCmd.Flags().String("environment-identifier", "", "The ID of the environment in which the service is created.")
	migrationHubRefactorSpaces_createServiceCmd.Flags().String("lambda-endpoint", "", "The configuration for the Lambda endpoint type.")
	migrationHubRefactorSpaces_createServiceCmd.Flags().String("name", "", "The name of the service.")
	migrationHubRefactorSpaces_createServiceCmd.Flags().String("tags", "", "The tags to assign to the service.")
	migrationHubRefactorSpaces_createServiceCmd.Flags().String("url-endpoint", "", "The configuration for the URL endpoint type.")
	migrationHubRefactorSpaces_createServiceCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
	migrationHubRefactorSpaces_createServiceCmd.MarkFlagRequired("application-identifier")
	migrationHubRefactorSpaces_createServiceCmd.MarkFlagRequired("endpoint-type")
	migrationHubRefactorSpaces_createServiceCmd.MarkFlagRequired("environment-identifier")
	migrationHubRefactorSpaces_createServiceCmd.MarkFlagRequired("name")
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_createServiceCmd)
}
