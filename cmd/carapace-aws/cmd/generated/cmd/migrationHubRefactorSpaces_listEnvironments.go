package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_listEnvironmentsCmd = &cobra.Command{
	Use:   "list-environments",
	Short: "Lists Amazon Web Services Migration Hub Refactor Spaces environments owned by a caller account or shared with the caller account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_listEnvironmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpaces_listEnvironmentsCmd).Standalone()

		migrationHubRefactorSpaces_listEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		migrationHubRefactorSpaces_listEnvironmentsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	})
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_listEnvironmentsCmd)
}
