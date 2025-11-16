package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_stopWorkflowRunCmd = &cobra.Command{
	Use:   "stop-workflow-run",
	Short: "Stops the execution of the specified workflow run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_stopWorkflowRunCmd).Standalone()

	glue_stopWorkflowRunCmd.Flags().String("name", "", "The name of the workflow to stop.")
	glue_stopWorkflowRunCmd.Flags().String("run-id", "", "The ID of the workflow run to stop.")
	glue_stopWorkflowRunCmd.MarkFlagRequired("name")
	glue_stopWorkflowRunCmd.MarkFlagRequired("run-id")
	glueCmd.AddCommand(glue_stopWorkflowRunCmd)
}
