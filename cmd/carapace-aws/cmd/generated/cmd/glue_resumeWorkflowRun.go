package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_resumeWorkflowRunCmd = &cobra.Command{
	Use:   "resume-workflow-run",
	Short: "Restarts selected nodes of a previous partially completed workflow run and resumes the workflow run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_resumeWorkflowRunCmd).Standalone()

	glue_resumeWorkflowRunCmd.Flags().String("name", "", "The name of the workflow to resume.")
	glue_resumeWorkflowRunCmd.Flags().String("node-ids", "", "A list of the node IDs for the nodes you want to restart.")
	glue_resumeWorkflowRunCmd.Flags().String("run-id", "", "The ID of the workflow run to resume.")
	glue_resumeWorkflowRunCmd.MarkFlagRequired("name")
	glue_resumeWorkflowRunCmd.MarkFlagRequired("node-ids")
	glue_resumeWorkflowRunCmd.MarkFlagRequired("run-id")
	glueCmd.AddCommand(glue_resumeWorkflowRunCmd)
}
