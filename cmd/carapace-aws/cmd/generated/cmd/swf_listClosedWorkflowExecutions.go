package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_listClosedWorkflowExecutionsCmd = &cobra.Command{
	Use:   "list-closed-workflow-executions",
	Short: "Returns a list of closed workflow executions in the specified domain that meet the filtering criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_listClosedWorkflowExecutionsCmd).Standalone()

	swf_listClosedWorkflowExecutionsCmd.Flags().String("close-status-filter", "", "If specified, only workflow executions that match this *close status* are listed.")
	swf_listClosedWorkflowExecutionsCmd.Flags().String("close-time-filter", "", "If specified, the workflow executions are included in the returned results based on whether their close times are within the range specified by this filter.")
	swf_listClosedWorkflowExecutionsCmd.Flags().String("domain", "", "The name of the domain that contains the workflow executions to list.")
	swf_listClosedWorkflowExecutionsCmd.Flags().String("execution-filter", "", "If specified, only workflow executions matching the workflow ID specified in the filter are returned.")
	swf_listClosedWorkflowExecutionsCmd.Flags().String("maximum-page-size", "", "The maximum number of results that are returned per call.")
	swf_listClosedWorkflowExecutionsCmd.Flags().String("next-page-token", "", "If `NextPageToken` is returned there are more results available.")
	swf_listClosedWorkflowExecutionsCmd.Flags().String("reverse-order", "", "When set to `true`, returns the results in reverse order.")
	swf_listClosedWorkflowExecutionsCmd.Flags().String("start-time-filter", "", "If specified, the workflow executions are included in the returned results based on whether their start times are within the range specified by this filter.")
	swf_listClosedWorkflowExecutionsCmd.Flags().String("tag-filter", "", "If specified, only executions that have the matching tag are listed.")
	swf_listClosedWorkflowExecutionsCmd.Flags().String("type-filter", "", "If specified, only executions of the type specified in the filter are returned.")
	swf_listClosedWorkflowExecutionsCmd.MarkFlagRequired("domain")
	swfCmd.AddCommand(swf_listClosedWorkflowExecutionsCmd)
}
