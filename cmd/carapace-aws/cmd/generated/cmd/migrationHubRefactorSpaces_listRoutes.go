package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_listRoutesCmd = &cobra.Command{
	Use:   "list-routes",
	Short: "Lists all the Amazon Web Services Migration Hub Refactor Spaces routes within an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_listRoutesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpaces_listRoutesCmd).Standalone()

		migrationHubRefactorSpaces_listRoutesCmd.Flags().String("application-identifier", "", "The ID of the application.")
		migrationHubRefactorSpaces_listRoutesCmd.Flags().String("environment-identifier", "", "The ID of the environment.")
		migrationHubRefactorSpaces_listRoutesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		migrationHubRefactorSpaces_listRoutesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		migrationHubRefactorSpaces_listRoutesCmd.MarkFlagRequired("application-identifier")
		migrationHubRefactorSpaces_listRoutesCmd.MarkFlagRequired("environment-identifier")
	})
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_listRoutesCmd)
}
