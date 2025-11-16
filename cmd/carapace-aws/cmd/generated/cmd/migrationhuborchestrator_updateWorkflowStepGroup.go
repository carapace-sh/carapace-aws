package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_updateWorkflowStepGroupCmd = &cobra.Command{
	Use:   "update-workflow-step-group",
	Short: "Update the step group in a migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_updateWorkflowStepGroupCmd).Standalone()

	migrationhuborchestrator_updateWorkflowStepGroupCmd.Flags().String("description", "", "The description of the step group.")
	migrationhuborchestrator_updateWorkflowStepGroupCmd.Flags().String("id", "", "The ID of the step group.")
	migrationhuborchestrator_updateWorkflowStepGroupCmd.Flags().String("name", "", "The name of the step group.")
	migrationhuborchestrator_updateWorkflowStepGroupCmd.Flags().String("next", "", "The next step group.")
	migrationhuborchestrator_updateWorkflowStepGroupCmd.Flags().String("previous", "", "The previous step group.")
	migrationhuborchestrator_updateWorkflowStepGroupCmd.Flags().String("workflow-id", "", "The ID of the migration workflow.")
	migrationhuborchestrator_updateWorkflowStepGroupCmd.MarkFlagRequired("id")
	migrationhuborchestrator_updateWorkflowStepGroupCmd.MarkFlagRequired("workflow-id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_updateWorkflowStepGroupCmd)
}
