package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_listTemplateStepsCmd = &cobra.Command{
	Use:   "list-template-steps",
	Short: "List the steps in a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_listTemplateStepsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhuborchestrator_listTemplateStepsCmd).Standalone()

		migrationhuborchestrator_listTemplateStepsCmd.Flags().String("max-results", "", "The maximum number of results that can be returned.")
		migrationhuborchestrator_listTemplateStepsCmd.Flags().String("next-token", "", "The pagination token.")
		migrationhuborchestrator_listTemplateStepsCmd.Flags().String("step-group-id", "", "The ID of the step group.")
		migrationhuborchestrator_listTemplateStepsCmd.Flags().String("template-id", "", "The ID of the template.")
		migrationhuborchestrator_listTemplateStepsCmd.MarkFlagRequired("step-group-id")
		migrationhuborchestrator_listTemplateStepsCmd.MarkFlagRequired("template-id")
	})
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_listTemplateStepsCmd)
}
