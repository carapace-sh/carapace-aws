package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateWorkflowCmd = &cobra.Command{
	Use:   "update-workflow",
	Short: "Updates an existing workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_updateWorkflowCmd).Standalone()

		glue_updateWorkflowCmd.Flags().String("default-run-properties", "", "A collection of properties to be used as part of each execution of the workflow.")
		glue_updateWorkflowCmd.Flags().String("description", "", "The description of the workflow.")
		glue_updateWorkflowCmd.Flags().String("max-concurrent-runs", "", "You can use this parameter to prevent unwanted multiple updates to data, to control costs, or in some cases, to prevent exceeding the maximum number of concurrent runs of any of the component jobs.")
		glue_updateWorkflowCmd.Flags().String("name", "", "Name of the workflow to be updated.")
		glue_updateWorkflowCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_updateWorkflowCmd)
}
