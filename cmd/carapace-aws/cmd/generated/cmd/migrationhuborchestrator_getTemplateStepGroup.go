package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_getTemplateStepGroupCmd = &cobra.Command{
	Use:   "get-template-step-group",
	Short: "Get a step group in a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_getTemplateStepGroupCmd).Standalone()

	migrationhuborchestrator_getTemplateStepGroupCmd.Flags().String("id", "", "The ID of the step group.")
	migrationhuborchestrator_getTemplateStepGroupCmd.Flags().String("template-id", "", "The ID of the template.")
	migrationhuborchestrator_getTemplateStepGroupCmd.MarkFlagRequired("id")
	migrationhuborchestrator_getTemplateStepGroupCmd.MarkFlagRequired("template-id")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_getTemplateStepGroupCmd)
}
