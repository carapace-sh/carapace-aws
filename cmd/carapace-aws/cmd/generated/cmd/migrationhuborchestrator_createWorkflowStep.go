package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_createWorkflowStepCmd = &cobra.Command{
	Use:   "create-workflow-step",
	Short: "Create a step in the migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_createWorkflowStepCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhuborchestrator_createWorkflowStepCmd).Standalone()

		migrationhuborchestrator_createWorkflowStepCmd.Flags().String("description", "", "The description of the step.")
		migrationhuborchestrator_createWorkflowStepCmd.Flags().String("name", "", "The name of the step.")
		migrationhuborchestrator_createWorkflowStepCmd.Flags().String("next", "", "The next step.")
		migrationhuborchestrator_createWorkflowStepCmd.Flags().String("outputs", "", "The key value pairs added for the expected output.")
		migrationhuborchestrator_createWorkflowStepCmd.Flags().String("previous", "", "The previous step.")
		migrationhuborchestrator_createWorkflowStepCmd.Flags().String("step-action-type", "", "The action type of the step.")
		migrationhuborchestrator_createWorkflowStepCmd.Flags().String("step-group-id", "", "The ID of the step group.")
		migrationhuborchestrator_createWorkflowStepCmd.Flags().String("step-target", "", "The servers on which a step will be run.")
		migrationhuborchestrator_createWorkflowStepCmd.Flags().String("workflow-id", "", "The ID of the migration workflow.")
		migrationhuborchestrator_createWorkflowStepCmd.Flags().String("workflow-step-automation-configuration", "", "The custom script to run tests on source or target environments.")
		migrationhuborchestrator_createWorkflowStepCmd.MarkFlagRequired("name")
		migrationhuborchestrator_createWorkflowStepCmd.MarkFlagRequired("step-action-type")
		migrationhuborchestrator_createWorkflowStepCmd.MarkFlagRequired("step-group-id")
		migrationhuborchestrator_createWorkflowStepCmd.MarkFlagRequired("workflow-id")
	})
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_createWorkflowStepCmd)
}
