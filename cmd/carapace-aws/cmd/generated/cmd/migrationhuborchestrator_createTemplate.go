package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_createTemplateCmd = &cobra.Command{
	Use:   "create-template",
	Short: "Creates a migration workflow template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_createTemplateCmd).Standalone()

	migrationhuborchestrator_createTemplateCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	migrationhuborchestrator_createTemplateCmd.Flags().String("tags", "", "The tags to add to the migration workflow template.")
	migrationhuborchestrator_createTemplateCmd.Flags().String("template-description", "", "A description of the migration workflow template.")
	migrationhuborchestrator_createTemplateCmd.Flags().String("template-name", "", "The name of the migration workflow template.")
	migrationhuborchestrator_createTemplateCmd.Flags().String("template-source", "", "The source of the migration workflow template.")
	migrationhuborchestrator_createTemplateCmd.MarkFlagRequired("template-name")
	migrationhuborchestrator_createTemplateCmd.MarkFlagRequired("template-source")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_createTemplateCmd)
}
