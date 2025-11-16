package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_listPluginsCmd = &cobra.Command{
	Use:   "list-plugins",
	Short: "List AWS Migration Hub Orchestrator plugins.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_listPluginsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhuborchestrator_listPluginsCmd).Standalone()

		migrationhuborchestrator_listPluginsCmd.Flags().String("max-results", "", "The maximum number of plugins that can be returned.")
		migrationhuborchestrator_listPluginsCmd.Flags().String("next-token", "", "The pagination token.")
	})
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_listPluginsCmd)
}
