package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpacesCmd = &cobra.Command{
	Use:   "migration-hub-refactor-spaces",
	Short: "Amazon Web Services Migration Hub Refactor Spaces\n\nThis API reference provides descriptions, syntax, and other details about each of the actions and data types for Amazon Web Services Migration Hub Refactor Spaces (Refactor Spaces).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpacesCmd).Standalone()

	})
	rootCmd.AddCommand(migrationHubRefactorSpacesCmd)
}
