package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_getWorkflowRunCmd = &cobra.Command{
	Use:   "get-workflow-run",
	Short: "Returns information about a specified run of a workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_getWorkflowRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_getWorkflowRunCmd).Standalone()

		codecatalyst_getWorkflowRunCmd.Flags().String("id", "", "The ID of the workflow run.")
		codecatalyst_getWorkflowRunCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_getWorkflowRunCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_getWorkflowRunCmd.MarkFlagRequired("id")
		codecatalyst_getWorkflowRunCmd.MarkFlagRequired("project-name")
		codecatalyst_getWorkflowRunCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_getWorkflowRunCmd)
}
