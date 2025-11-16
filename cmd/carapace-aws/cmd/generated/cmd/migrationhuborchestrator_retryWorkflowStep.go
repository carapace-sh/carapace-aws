package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_retryWorkflowStepCmd = &cobra.Command{
	Use:   "retry-workflow-step",
	Short: "Retry a failed step in a migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_retryWorkflowStepCmd).Standalone()

	migrationhuborchestrator_retryWorkflowStepCmd.Flags().String("id", "", "The ID of the step.")
	migrationhuborchestrator_retryWorkflowStepCmd.Flags().String("step-group-id", "", "The ID of the step group.")
	migrationhuborchestrator_retryWorkflowStepCmd.Flags().String("workflow-id", "", "The ID of the migration workflow.")
	migrationhuborchestrator_retryWorkflowStepCmd.MarkFlagRequired("id")
	migrationhuborchestrator_retryWorkflowStepCmd.MarkFlagRequired("step-group-id")
	migrationhuborchestrator_retryWorkflowStepCmd.MarkFlagRequired("workflow-id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_retryWorkflowStepCmd)
}
