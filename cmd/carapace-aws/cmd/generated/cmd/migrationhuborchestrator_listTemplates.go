package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_listTemplatesCmd = &cobra.Command{
	Use:   "list-templates",
	Short: "List the templates available in Migration Hub Orchestrator to create a migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_listTemplatesCmd).Standalone()

	migrationhuborchestrator_listTemplatesCmd.Flags().String("max-results", "", "The maximum number of results that can be returned.")
	migrationhuborchestrator_listTemplatesCmd.Flags().String("name", "", "The name of the template.")
	migrationhuborchestrator_listTemplatesCmd.Flags().String("next-token", "", "The pagination token.")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_listTemplatesCmd)
}
