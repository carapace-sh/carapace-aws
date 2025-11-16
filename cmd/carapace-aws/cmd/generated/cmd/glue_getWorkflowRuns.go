package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getWorkflowRunsCmd = &cobra.Command{
	Use:   "get-workflow-runs",
	Short: "Retrieves metadata for all runs of a given workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getWorkflowRunsCmd).Standalone()

	glue_getWorkflowRunsCmd.Flags().String("include-graph", "", "Specifies whether to include the workflow graph in response or not.")
	glue_getWorkflowRunsCmd.Flags().String("max-results", "", "The maximum number of workflow runs to be included in the response.")
	glue_getWorkflowRunsCmd.Flags().String("name", "", "Name of the workflow whose metadata of runs should be returned.")
	glue_getWorkflowRunsCmd.Flags().String("next-token", "", "The maximum size of the response.")
	glue_getWorkflowRunsCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_getWorkflowRunsCmd)
}
