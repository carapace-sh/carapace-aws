package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_countClosedWorkflowExecutionsCmd = &cobra.Command{
	Use:   "count-closed-workflow-executions",
	Short: "Returns the number of closed workflow executions within the given domain that meet the specified filtering criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_countClosedWorkflowExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_countClosedWorkflowExecutionsCmd).Standalone()

		swf_countClosedWorkflowExecutionsCmd.Flags().String("close-status-filter", "", "If specified, only workflow executions that match this close status are counted.")
		swf_countClosedWorkflowExecutionsCmd.Flags().String("close-time-filter", "", "If specified, only workflow executions that meet the close time criteria of the filter are counted.")
		swf_countClosedWorkflowExecutionsCmd.Flags().String("domain", "", "The name of the domain containing the workflow executions to count.")
		swf_countClosedWorkflowExecutionsCmd.Flags().String("execution-filter", "", "If specified, only workflow executions matching the `WorkflowId` in the filter are counted.")
		swf_countClosedWorkflowExecutionsCmd.Flags().String("start-time-filter", "", "If specified, only workflow executions that meet the start time criteria of the filter are counted.")
		swf_countClosedWorkflowExecutionsCmd.Flags().String("tag-filter", "", "If specified, only executions that have a tag that matches the filter are counted.")
		swf_countClosedWorkflowExecutionsCmd.Flags().String("type-filter", "", "If specified, indicates the type of the workflow executions to be counted.")
		swf_countClosedWorkflowExecutionsCmd.MarkFlagRequired("domain")
	})
	swfCmd.AddCommand(swf_countClosedWorkflowExecutionsCmd)
}
