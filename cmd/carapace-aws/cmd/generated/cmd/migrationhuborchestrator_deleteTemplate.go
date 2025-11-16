package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_deleteTemplateCmd = &cobra.Command{
	Use:   "delete-template",
	Short: "Deletes a migration workflow template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_deleteTemplateCmd).Standalone()

	migrationhuborchestrator_deleteTemplateCmd.Flags().String("id", "", "The ID of the request to delete a migration workflow template.")
	migrationhuborchestrator_deleteTemplateCmd.MarkFlagRequired("id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_deleteTemplateCmd)
}
