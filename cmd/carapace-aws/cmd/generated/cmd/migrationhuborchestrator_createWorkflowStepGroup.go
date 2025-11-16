package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_createWorkflowStepGroupCmd = &cobra.Command{
	Use:   "create-workflow-step-group",
	Short: "Create a step group in a migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_createWorkflowStepGroupCmd).Standalone()

	migrationhuborchestrator_createWorkflowStepGroupCmd.Flags().String("description", "", "The description of the step group.")
	migrationhuborchestrator_createWorkflowStepGroupCmd.Flags().String("name", "", "The name of the step group.")
	migrationhuborchestrator_createWorkflowStepGroupCmd.Flags().String("next", "", "The next step group.")
	migrationhuborchestrator_createWorkflowStepGroupCmd.Flags().String("previous", "", "The previous step group.")
	migrationhuborchestrator_createWorkflowStepGroupCmd.Flags().String("workflow-id", "", "The ID of the migration workflow that will contain the step group.")
	migrationhuborchestrator_createWorkflowStepGroupCmd.MarkFlagRequired("name")
	migrationhuborchestrator_createWorkflowStepGroupCmd.MarkFlagRequired("workflow-id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_createWorkflowStepGroupCmd)
}
