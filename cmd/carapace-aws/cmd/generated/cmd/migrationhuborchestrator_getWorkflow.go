package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_getWorkflowCmd = &cobra.Command{
	Use:   "get-workflow",
	Short: "Get migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_getWorkflowCmd).Standalone()

	migrationhuborchestrator_getWorkflowCmd.Flags().String("id", "", "The ID of the migration workflow.")
	migrationhuborchestrator_getWorkflowCmd.MarkFlagRequired("id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_getWorkflowCmd)
}
