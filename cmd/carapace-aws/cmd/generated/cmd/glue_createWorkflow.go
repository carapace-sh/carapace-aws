package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createWorkflowCmd = &cobra.Command{
	Use:   "create-workflow",
	Short: "Creates a new workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createWorkflowCmd).Standalone()

		glue_createWorkflowCmd.Flags().String("default-run-properties", "", "A collection of properties to be used as part of each execution of the workflow.")
		glue_createWorkflowCmd.Flags().String("description", "", "A description of the workflow.")
		glue_createWorkflowCmd.Flags().String("max-concurrent-runs", "", "You can use this parameter to prevent unwanted multiple updates to data, to control costs, or in some cases, to prevent exceeding the maximum number of concurrent runs of any of the component jobs.")
		glue_createWorkflowCmd.Flags().String("name", "", "The name to be assigned to the workflow.")
		glue_createWorkflowCmd.Flags().String("tags", "", "The tags to be used with this workflow.")
		glue_createWorkflowCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_createWorkflowCmd)
}
