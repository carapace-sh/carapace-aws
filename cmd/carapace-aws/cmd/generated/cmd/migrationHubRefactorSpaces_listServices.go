package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_listServicesCmd = &cobra.Command{
	Use:   "list-services",
	Short: "Lists all the Amazon Web Services Migration Hub Refactor Spaces services within an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_listServicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpaces_listServicesCmd).Standalone()

		migrationHubRefactorSpaces_listServicesCmd.Flags().String("application-identifier", "", "The ID of the application.")
		migrationHubRefactorSpaces_listServicesCmd.Flags().String("environment-identifier", "", "The ID of the environment.")
		migrationHubRefactorSpaces_listServicesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		migrationHubRefactorSpaces_listServicesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		migrationHubRefactorSpaces_listServicesCmd.MarkFlagRequired("application-identifier")
		migrationHubRefactorSpaces_listServicesCmd.MarkFlagRequired("environment-identifier")
	})
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_listServicesCmd)
}
