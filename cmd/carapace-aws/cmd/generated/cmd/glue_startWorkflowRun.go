package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_startWorkflowRunCmd = &cobra.Command{
	Use:   "start-workflow-run",
	Short: "Starts a new run of the specified workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_startWorkflowRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_startWorkflowRunCmd).Standalone()

		glue_startWorkflowRunCmd.Flags().String("name", "", "The name of the workflow to start.")
		glue_startWorkflowRunCmd.Flags().String("run-properties", "", "The workflow run properties for the new workflow run.")
		glue_startWorkflowRunCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_startWorkflowRunCmd)
}
