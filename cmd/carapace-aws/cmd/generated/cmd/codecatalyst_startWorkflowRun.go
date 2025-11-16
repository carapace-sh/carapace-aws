package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_startWorkflowRunCmd = &cobra.Command{
	Use:   "start-workflow-run",
	Short: "Begins a run of a specified workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_startWorkflowRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_startWorkflowRunCmd).Standalone()

		codecatalyst_startWorkflowRunCmd.Flags().String("client-token", "", "A user-specified idempotency token.")
		codecatalyst_startWorkflowRunCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_startWorkflowRunCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_startWorkflowRunCmd.Flags().String("workflow-id", "", "The system-generated unique ID of the workflow.")
		codecatalyst_startWorkflowRunCmd.MarkFlagRequired("project-name")
		codecatalyst_startWorkflowRunCmd.MarkFlagRequired("space-name")
		codecatalyst_startWorkflowRunCmd.MarkFlagRequired("workflow-id")
	})
	codecatalystCmd.AddCommand(codecatalyst_startWorkflowRunCmd)
}
