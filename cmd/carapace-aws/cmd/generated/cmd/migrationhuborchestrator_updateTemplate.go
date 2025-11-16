package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_updateTemplateCmd = &cobra.Command{
	Use:   "update-template",
	Short: "Updates a migration workflow template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_updateTemplateCmd).Standalone()

	migrationhuborchestrator_updateTemplateCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	migrationhuborchestrator_updateTemplateCmd.Flags().String("id", "", "The ID of the request to update a migration workflow template.")
	migrationhuborchestrator_updateTemplateCmd.Flags().String("template-description", "", "The description of the migration workflow template to update.")
	migrationhuborchestrator_updateTemplateCmd.Flags().String("template-name", "", "The name of the migration workflow template to update.")
	migrationhuborchestrator_updateTemplateCmd.MarkFlagRequired("id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_updateTemplateCmd)
}
