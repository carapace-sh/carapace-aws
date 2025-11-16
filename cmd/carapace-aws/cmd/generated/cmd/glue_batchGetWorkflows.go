package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchGetWorkflowsCmd = &cobra.Command{
	Use:   "batch-get-workflows",
	Short: "Returns a list of resource metadata for a given list of workflow names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchGetWorkflowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_batchGetWorkflowsCmd).Standalone()

		glue_batchGetWorkflowsCmd.Flags().String("include-graph", "", "Specifies whether to include a graph when returning the workflow resource metadata.")
		glue_batchGetWorkflowsCmd.Flags().String("names", "", "A list of workflow names, which may be the names returned from the `ListWorkflows` operation.")
		glue_batchGetWorkflowsCmd.MarkFlagRequired("names")
	})
	glueCmd.AddCommand(glue_batchGetWorkflowsCmd)
}
