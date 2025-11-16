package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getWorkflowCmd = &cobra.Command{
	Use:   "get-workflow",
	Short: "Retrieves resource metadata for a workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getWorkflowCmd).Standalone()

		glue_getWorkflowCmd.Flags().String("include-graph", "", "Specifies whether to include a graph when returning the workflow resource metadata.")
		glue_getWorkflowCmd.Flags().String("name", "", "The name of the workflow to retrieve.")
		glue_getWorkflowCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_getWorkflowCmd)
}
