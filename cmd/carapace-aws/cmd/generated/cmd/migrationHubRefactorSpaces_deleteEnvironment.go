package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_deleteEnvironmentCmd = &cobra.Command{
	Use:   "delete-environment",
	Short: "Deletes an Amazon Web Services Migration Hub Refactor Spaces environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_deleteEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpaces_deleteEnvironmentCmd).Standalone()

		migrationHubRefactorSpaces_deleteEnvironmentCmd.Flags().String("environment-identifier", "", "The ID of the environment.")
		migrationHubRefactorSpaces_deleteEnvironmentCmd.MarkFlagRequired("environment-identifier")
	})
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_deleteEnvironmentCmd)
}
