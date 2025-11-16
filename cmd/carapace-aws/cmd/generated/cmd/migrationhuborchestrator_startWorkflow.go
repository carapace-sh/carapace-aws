package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_startWorkflowCmd = &cobra.Command{
	Use:   "start-workflow",
	Short: "Start a migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_startWorkflowCmd).Standalone()

	migrationhuborchestrator_startWorkflowCmd.Flags().String("id", "", "The ID of the migration workflow.")
	migrationhuborchestrator_startWorkflowCmd.MarkFlagRequired("id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_startWorkflowCmd)
}
