package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_getWorkflowStepCmd = &cobra.Command{
	Use:   "get-workflow-step",
	Short: "Get a step in the migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_getWorkflowStepCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhuborchestrator_getWorkflowStepCmd).Standalone()

		migrationhuborchestrator_getWorkflowStepCmd.Flags().String("id", "", "The ID of the step.")
		migrationhuborchestrator_getWorkflowStepCmd.Flags().String("step-group-id", "", "The ID of the step group.")
		migrationhuborchestrator_getWorkflowStepCmd.Flags().String("workflow-id", "", "The ID of the migration workflow.")
		migrationhuborchestrator_getWorkflowStepCmd.MarkFlagRequired("id")
		migrationhuborchestrator_getWorkflowStepCmd.MarkFlagRequired("step-group-id")
		migrationhuborchestrator_getWorkflowStepCmd.MarkFlagRequired("workflow-id")
	})
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_getWorkflowStepCmd)
}
