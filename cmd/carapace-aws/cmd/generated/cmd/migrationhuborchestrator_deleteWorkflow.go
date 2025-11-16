package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_deleteWorkflowCmd = &cobra.Command{
	Use:   "delete-workflow",
	Short: "Delete a migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_deleteWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhuborchestrator_deleteWorkflowCmd).Standalone()

		migrationhuborchestrator_deleteWorkflowCmd.Flags().String("id", "", "The ID of the migration workflow you want to delete.")
		migrationhuborchestrator_deleteWorkflowCmd.MarkFlagRequired("id")
	})
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_deleteWorkflowCmd)
}
