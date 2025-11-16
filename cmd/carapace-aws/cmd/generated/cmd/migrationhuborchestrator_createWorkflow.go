package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_createWorkflowCmd = &cobra.Command{
	Use:   "create-workflow",
	Short: "Create a workflow to orchestrate your migrations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_createWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhuborchestrator_createWorkflowCmd).Standalone()

		migrationhuborchestrator_createWorkflowCmd.Flags().String("application-configuration-id", "", "The configuration ID of the application configured in Application Discovery Service.")
		migrationhuborchestrator_createWorkflowCmd.Flags().String("description", "", "The description of the migration workflow.")
		migrationhuborchestrator_createWorkflowCmd.Flags().String("input-parameters", "", "The input parameters required to create a migration workflow.")
		migrationhuborchestrator_createWorkflowCmd.Flags().String("name", "", "The name of the migration workflow.")
		migrationhuborchestrator_createWorkflowCmd.Flags().String("step-targets", "", "The servers on which a step will be run.")
		migrationhuborchestrator_createWorkflowCmd.Flags().String("tags", "", "The tags to add on a migration workflow.")
		migrationhuborchestrator_createWorkflowCmd.Flags().String("template-id", "", "The ID of the template.")
		migrationhuborchestrator_createWorkflowCmd.MarkFlagRequired("input-parameters")
		migrationhuborchestrator_createWorkflowCmd.MarkFlagRequired("name")
		migrationhuborchestrator_createWorkflowCmd.MarkFlagRequired("template-id")
	})
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_createWorkflowCmd)
}
