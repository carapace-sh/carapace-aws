package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_listWorkflowStepGroupsCmd = &cobra.Command{
	Use:   "list-workflow-step-groups",
	Short: "List the step groups in a migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_listWorkflowStepGroupsCmd).Standalone()

	migrationhuborchestrator_listWorkflowStepGroupsCmd.Flags().String("max-results", "", "The maximum number of results that can be returned.")
	migrationhuborchestrator_listWorkflowStepGroupsCmd.Flags().String("next-token", "", "The pagination token.")
	migrationhuborchestrator_listWorkflowStepGroupsCmd.Flags().String("workflow-id", "", "The ID of the migration workflow.")
	migrationhuborchestrator_listWorkflowStepGroupsCmd.MarkFlagRequired("workflow-id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_listWorkflowStepGroupsCmd)
}
