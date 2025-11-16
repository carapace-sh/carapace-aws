package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_updateWorkflowStepCmd = &cobra.Command{
	Use:   "update-workflow-step",
	Short: "Update a step in a migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_updateWorkflowStepCmd).Standalone()

	migrationhuborchestrator_updateWorkflowStepCmd.Flags().String("description", "", "The description of the step.")
	migrationhuborchestrator_updateWorkflowStepCmd.Flags().String("id", "", "The ID of the step.")
	migrationhuborchestrator_updateWorkflowStepCmd.Flags().String("name", "", "The name of the step.")
	migrationhuborchestrator_updateWorkflowStepCmd.Flags().String("next", "", "The next step.")
	migrationhuborchestrator_updateWorkflowStepCmd.Flags().String("outputs", "", "The outputs of a step.")
	migrationhuborchestrator_updateWorkflowStepCmd.Flags().String("previous", "", "The previous step.")
	migrationhuborchestrator_updateWorkflowStepCmd.Flags().String("status", "", "The status of the step.")
	migrationhuborchestrator_updateWorkflowStepCmd.Flags().String("step-action-type", "", "The action type of the step.")
	migrationhuborchestrator_updateWorkflowStepCmd.Flags().String("step-group-id", "", "The ID of the step group.")
	migrationhuborchestrator_updateWorkflowStepCmd.Flags().String("step-target", "", "The servers on which a step will be run.")
	migrationhuborchestrator_updateWorkflowStepCmd.Flags().String("workflow-id", "", "The ID of the migration workflow.")
	migrationhuborchestrator_updateWorkflowStepCmd.Flags().String("workflow-step-automation-configuration", "", "The custom script to run tests on the source and target environments.")
	migrationhuborchestrator_updateWorkflowStepCmd.MarkFlagRequired("id")
	migrationhuborchestrator_updateWorkflowStepCmd.MarkFlagRequired("step-group-id")
	migrationhuborchestrator_updateWorkflowStepCmd.MarkFlagRequired("workflow-id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_updateWorkflowStepCmd)
}
