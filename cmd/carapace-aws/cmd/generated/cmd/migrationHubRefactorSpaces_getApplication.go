package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_getApplicationCmd = &cobra.Command{
	Use:   "get-application",
	Short: "Gets an Amazon Web Services Migration Hub Refactor Spaces application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_getApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpaces_getApplicationCmd).Standalone()

		migrationHubRefactorSpaces_getApplicationCmd.Flags().String("application-identifier", "", "The ID of the application.")
		migrationHubRefactorSpaces_getApplicationCmd.Flags().String("environment-identifier", "", "The ID of the environment.")
		migrationHubRefactorSpaces_getApplicationCmd.MarkFlagRequired("application-identifier")
		migrationHubRefactorSpaces_getApplicationCmd.MarkFlagRequired("environment-identifier")
	})
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_getApplicationCmd)
}
