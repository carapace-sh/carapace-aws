package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Deletes an Amazon Web Services Migration Hub Refactor Spaces application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_deleteApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpaces_deleteApplicationCmd).Standalone()

		migrationHubRefactorSpaces_deleteApplicationCmd.Flags().String("application-identifier", "", "The ID of the application.")
		migrationHubRefactorSpaces_deleteApplicationCmd.Flags().String("environment-identifier", "", "The ID of the environment.")
		migrationHubRefactorSpaces_deleteApplicationCmd.MarkFlagRequired("application-identifier")
		migrationHubRefactorSpaces_deleteApplicationCmd.MarkFlagRequired("environment-identifier")
	})
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_deleteApplicationCmd)
}
