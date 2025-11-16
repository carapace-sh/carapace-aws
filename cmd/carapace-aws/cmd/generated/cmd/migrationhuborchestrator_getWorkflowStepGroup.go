package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_getWorkflowStepGroupCmd = &cobra.Command{
	Use:   "get-workflow-step-group",
	Short: "Get the step group of a migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_getWorkflowStepGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhuborchestrator_getWorkflowStepGroupCmd).Standalone()

		migrationhuborchestrator_getWorkflowStepGroupCmd.Flags().String("id", "", "The ID of the step group.")
		migrationhuborchestrator_getWorkflowStepGroupCmd.Flags().String("workflow-id", "", "The ID of the migration workflow.")
		migrationhuborchestrator_getWorkflowStepGroupCmd.MarkFlagRequired("id")
		migrationhuborchestrator_getWorkflowStepGroupCmd.MarkFlagRequired("workflow-id")
	})
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_getWorkflowStepGroupCmd)
}
