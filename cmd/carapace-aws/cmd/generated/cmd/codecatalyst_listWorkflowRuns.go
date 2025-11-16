package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_listWorkflowRunsCmd = &cobra.Command{
	Use:   "list-workflow-runs",
	Short: "Retrieves a list of workflow runs of a specified workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_listWorkflowRunsCmd).Standalone()

	codecatalyst_listWorkflowRunsCmd.Flags().String("max-results", "", "The maximum number of results to show in a single call to this API.")
	codecatalyst_listWorkflowRunsCmd.Flags().String("next-token", "", "A token returned from a call to this API to indicate the next batch of results to return, if any.")
	codecatalyst_listWorkflowRunsCmd.Flags().String("project-name", "", "The name of the project in the space.")
	codecatalyst_listWorkflowRunsCmd.Flags().String("sort-by", "", "Information used to sort the items in the returned list.")
	codecatalyst_listWorkflowRunsCmd.Flags().String("space-name", "", "The name of the space.")
	codecatalyst_listWorkflowRunsCmd.Flags().String("workflow-id", "", "The ID of the workflow.")
	codecatalyst_listWorkflowRunsCmd.MarkFlagRequired("project-name")
	codecatalyst_listWorkflowRunsCmd.MarkFlagRequired("space-name")
	codecatalystCmd.AddCommand(codecatalyst_listWorkflowRunsCmd)
}
