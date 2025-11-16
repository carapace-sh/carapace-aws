package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_updateWorkflowCmd = &cobra.Command{
	Use:   "update-workflow",
	Short: "Update a migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_updateWorkflowCmd).Standalone()

	migrationhuborchestrator_updateWorkflowCmd.Flags().String("description", "", "The description of the migration workflow.")
	migrationhuborchestrator_updateWorkflowCmd.Flags().String("id", "", "The ID of the migration workflow.")
	migrationhuborchestrator_updateWorkflowCmd.Flags().String("input-parameters", "", "The input parameters required to update a migration workflow.")
	migrationhuborchestrator_updateWorkflowCmd.Flags().String("name", "", "The name of the migration workflow.")
	migrationhuborchestrator_updateWorkflowCmd.Flags().String("step-targets", "", "The servers on which a step will be run.")
	migrationhuborchestrator_updateWorkflowCmd.MarkFlagRequired("id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_updateWorkflowCmd)
}
