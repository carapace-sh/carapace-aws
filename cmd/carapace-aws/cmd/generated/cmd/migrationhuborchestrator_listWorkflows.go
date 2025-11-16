package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_listWorkflowsCmd = &cobra.Command{
	Use:   "list-workflows",
	Short: "List the migration workflows.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_listWorkflowsCmd).Standalone()

	migrationhuborchestrator_listWorkflowsCmd.Flags().String("ads-application-configuration-name", "", "The name of the application configured in Application Discovery Service.")
	migrationhuborchestrator_listWorkflowsCmd.Flags().String("max-results", "", "The maximum number of results that can be returned.")
	migrationhuborchestrator_listWorkflowsCmd.Flags().String("name", "", "The name of the migration workflow.")
	migrationhuborchestrator_listWorkflowsCmd.Flags().String("next-token", "", "The pagination token.")
	migrationhuborchestrator_listWorkflowsCmd.Flags().String("status", "", "The status of the migration workflow.")
	migrationhuborchestrator_listWorkflowsCmd.Flags().String("template-id", "", "The ID of the template.")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_listWorkflowsCmd)
}
