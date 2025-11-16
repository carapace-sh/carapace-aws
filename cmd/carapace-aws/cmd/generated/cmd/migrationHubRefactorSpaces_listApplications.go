package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Lists all the Amazon Web Services Migration Hub Refactor Spaces applications within an environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_listApplicationsCmd).Standalone()

	migrationHubRefactorSpaces_listApplicationsCmd.Flags().String("environment-identifier", "", "The ID of the environment.")
	migrationHubRefactorSpaces_listApplicationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	migrationHubRefactorSpaces_listApplicationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	migrationHubRefactorSpaces_listApplicationsCmd.MarkFlagRequired("environment-identifier")
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_listApplicationsCmd)
}
