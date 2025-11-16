package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listWorkflowVersionsCmd = &cobra.Command{
	Use:   "list-workflow-versions",
	Short: "Lists the workflow versions for the specified workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listWorkflowVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_listWorkflowVersionsCmd).Standalone()

		omics_listWorkflowVersionsCmd.Flags().String("max-results", "", "The maximum number of workflows to return in one page of results.")
		omics_listWorkflowVersionsCmd.Flags().String("starting-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		omics_listWorkflowVersionsCmd.Flags().String("type", "", "The workflow type.")
		omics_listWorkflowVersionsCmd.Flags().String("workflow-id", "", "The workflow's ID.")
		omics_listWorkflowVersionsCmd.Flags().String("workflow-owner-id", "", "The 12-digit account ID of the workflow owner.")
		omics_listWorkflowVersionsCmd.MarkFlagRequired("workflow-id")
	})
	omicsCmd.AddCommand(omics_listWorkflowVersionsCmd)
}
