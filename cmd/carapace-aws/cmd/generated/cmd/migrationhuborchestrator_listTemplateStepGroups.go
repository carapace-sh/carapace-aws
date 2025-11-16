package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_listTemplateStepGroupsCmd = &cobra.Command{
	Use:   "list-template-step-groups",
	Short: "List the step groups in a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_listTemplateStepGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhuborchestrator_listTemplateStepGroupsCmd).Standalone()

		migrationhuborchestrator_listTemplateStepGroupsCmd.Flags().String("max-results", "", "The maximum number of results that can be returned.")
		migrationhuborchestrator_listTemplateStepGroupsCmd.Flags().String("next-token", "", "The pagination token.")
		migrationhuborchestrator_listTemplateStepGroupsCmd.Flags().String("template-id", "", "The ID of the template.")
		migrationhuborchestrator_listTemplateStepGroupsCmd.MarkFlagRequired("template-id")
	})
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_listTemplateStepGroupsCmd)
}
