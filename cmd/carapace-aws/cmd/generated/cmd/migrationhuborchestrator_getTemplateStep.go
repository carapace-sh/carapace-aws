package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_getTemplateStepCmd = &cobra.Command{
	Use:   "get-template-step",
	Short: "Get a specific step in a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_getTemplateStepCmd).Standalone()

	migrationhuborchestrator_getTemplateStepCmd.Flags().String("id", "", "The ID of the step.")
	migrationhuborchestrator_getTemplateStepCmd.Flags().String("step-group-id", "", "The ID of the step group.")
	migrationhuborchestrator_getTemplateStepCmd.Flags().String("template-id", "", "The ID of the template.")
	migrationhuborchestrator_getTemplateStepCmd.MarkFlagRequired("id")
	migrationhuborchestrator_getTemplateStepCmd.MarkFlagRequired("step-group-id")
	migrationhuborchestrator_getTemplateStepCmd.MarkFlagRequired("template-id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_getTemplateStepCmd)
}
