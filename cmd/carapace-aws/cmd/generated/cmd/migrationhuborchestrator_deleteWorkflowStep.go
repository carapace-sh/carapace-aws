package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_deleteWorkflowStepCmd = &cobra.Command{
	Use:   "delete-workflow-step",
	Short: "Delete a step in a migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_deleteWorkflowStepCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhuborchestrator_deleteWorkflowStepCmd).Standalone()

		migrationhuborchestrator_deleteWorkflowStepCmd.Flags().String("id", "", "The ID of the step you want to delete.")
		migrationhuborchestrator_deleteWorkflowStepCmd.Flags().String("step-group-id", "", "The ID of the step group that contains the step you want to delete.")
		migrationhuborchestrator_deleteWorkflowStepCmd.Flags().String("workflow-id", "", "The ID of the migration workflow.")
		migrationhuborchestrator_deleteWorkflowStepCmd.MarkFlagRequired("id")
		migrationhuborchestrator_deleteWorkflowStepCmd.MarkFlagRequired("step-group-id")
		migrationhuborchestrator_deleteWorkflowStepCmd.MarkFlagRequired("workflow-id")
	})
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_deleteWorkflowStepCmd)
}
