package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_deleteWorkflowStepGroupCmd = &cobra.Command{
	Use:   "delete-workflow-step-group",
	Short: "Delete a step group in a migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_deleteWorkflowStepGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhuborchestrator_deleteWorkflowStepGroupCmd).Standalone()

		migrationhuborchestrator_deleteWorkflowStepGroupCmd.Flags().String("id", "", "The ID of the step group you want to delete.")
		migrationhuborchestrator_deleteWorkflowStepGroupCmd.Flags().String("workflow-id", "", "The ID of the migration workflow.")
		migrationhuborchestrator_deleteWorkflowStepGroupCmd.MarkFlagRequired("id")
		migrationhuborchestrator_deleteWorkflowStepGroupCmd.MarkFlagRequired("workflow-id")
	})
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_deleteWorkflowStepGroupCmd)
}
