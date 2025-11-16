package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_listWorkflowStepsCmd = &cobra.Command{
	Use:   "list-workflow-steps",
	Short: "List the steps in a workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_listWorkflowStepsCmd).Standalone()

	migrationhuborchestrator_listWorkflowStepsCmd.Flags().String("max-results", "", "The maximum number of results that can be returned.")
	migrationhuborchestrator_listWorkflowStepsCmd.Flags().String("next-token", "", "The pagination token.")
	migrationhuborchestrator_listWorkflowStepsCmd.Flags().String("step-group-id", "", "The ID of the step group.")
	migrationhuborchestrator_listWorkflowStepsCmd.Flags().String("workflow-id", "", "The ID of the migration workflow.")
	migrationhuborchestrator_listWorkflowStepsCmd.MarkFlagRequired("step-group-id")
	migrationhuborchestrator_listWorkflowStepsCmd.MarkFlagRequired("workflow-id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_listWorkflowStepsCmd)
}
