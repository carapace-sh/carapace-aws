package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getWorkflowRunCmd = &cobra.Command{
	Use:   "get-workflow-run",
	Short: "Retrieves the metadata for a given workflow run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getWorkflowRunCmd).Standalone()

	glue_getWorkflowRunCmd.Flags().String("include-graph", "", "Specifies whether to include the workflow graph in response or not.")
	glue_getWorkflowRunCmd.Flags().String("name", "", "Name of the workflow being run.")
	glue_getWorkflowRunCmd.Flags().String("run-id", "", "The ID of the workflow run.")
	glue_getWorkflowRunCmd.MarkFlagRequired("name")
	glue_getWorkflowRunCmd.MarkFlagRequired("run-id")
	glueCmd.AddCommand(glue_getWorkflowRunCmd)
}
