package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_stopWorkflowCmd = &cobra.Command{
	Use:   "stop-workflow",
	Short: "Stop an ongoing migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_stopWorkflowCmd).Standalone()

	migrationhuborchestrator_stopWorkflowCmd.Flags().String("id", "", "The ID of the migration workflow.")
	migrationhuborchestrator_stopWorkflowCmd.MarkFlagRequired("id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_stopWorkflowCmd)
}
