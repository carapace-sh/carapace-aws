package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_getTemplateCmd = &cobra.Command{
	Use:   "get-template",
	Short: "Get the template you want to use for creating a migration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_getTemplateCmd).Standalone()

	migrationhuborchestrator_getTemplateCmd.Flags().String("id", "", "The ID of the template.")
	migrationhuborchestrator_getTemplateCmd.MarkFlagRequired("id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_getTemplateCmd)
}
