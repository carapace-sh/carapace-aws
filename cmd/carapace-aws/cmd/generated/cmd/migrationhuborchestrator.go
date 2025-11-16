package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestratorCmd = &cobra.Command{
	Use:   "migrationhuborchestrator",
	Short: "This API reference provides descriptions, syntax, and other details about each of the actions and data types for AWS Migration Hub Orchestrator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestratorCmd).Standalone()

	rootCmd.AddCommand(migrationhuborchestratorCmd)
}
