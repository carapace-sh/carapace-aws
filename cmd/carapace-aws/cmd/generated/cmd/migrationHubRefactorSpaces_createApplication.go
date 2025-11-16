package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates an Amazon Web Services Migration Hub Refactor Spaces application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_createApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpaces_createApplicationCmd).Standalone()

		migrationHubRefactorSpaces_createApplicationCmd.Flags().String("api-gateway-proxy", "", "A wrapper object holding the API Gateway endpoint type and stage name for the proxy.")
		migrationHubRefactorSpaces_createApplicationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		migrationHubRefactorSpaces_createApplicationCmd.Flags().String("environment-identifier", "", "The unique identifier of the environment.")
		migrationHubRefactorSpaces_createApplicationCmd.Flags().String("name", "", "The name to use for the application.")
		migrationHubRefactorSpaces_createApplicationCmd.Flags().String("proxy-type", "", "The proxy type of the proxy created within the application.")
		migrationHubRefactorSpaces_createApplicationCmd.Flags().String("tags", "", "The tags to assign to the application.")
		migrationHubRefactorSpaces_createApplicationCmd.Flags().String("vpc-id", "", "The ID of the virtual private cloud (VPC).")
		migrationHubRefactorSpaces_createApplicationCmd.MarkFlagRequired("environment-identifier")
		migrationHubRefactorSpaces_createApplicationCmd.MarkFlagRequired("name")
		migrationHubRefactorSpaces_createApplicationCmd.MarkFlagRequired("proxy-type")
		migrationHubRefactorSpaces_createApplicationCmd.MarkFlagRequired("vpc-id")
	})
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_createApplicationCmd)
}
