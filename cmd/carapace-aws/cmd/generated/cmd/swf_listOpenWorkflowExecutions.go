package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_listOpenWorkflowExecutionsCmd = &cobra.Command{
	Use:   "list-open-workflow-executions",
	Short: "Returns a list of open workflow executions in the specified domain that meet the filtering criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_listOpenWorkflowExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_listOpenWorkflowExecutionsCmd).Standalone()

		swf_listOpenWorkflowExecutionsCmd.Flags().String("domain", "", "The name of the domain that contains the workflow executions to list.")
		swf_listOpenWorkflowExecutionsCmd.Flags().String("execution-filter", "", "If specified, only workflow executions matching the workflow ID specified in the filter are returned.")
		swf_listOpenWorkflowExecutionsCmd.Flags().String("maximum-page-size", "", "The maximum number of results that are returned per call.")
		swf_listOpenWorkflowExecutionsCmd.Flags().String("next-page-token", "", "If `NextPageToken` is returned there are more results available.")
		swf_listOpenWorkflowExecutionsCmd.Flags().String("reverse-order", "", "When set to `true`, returns the results in reverse order.")
		swf_listOpenWorkflowExecutionsCmd.Flags().String("start-time-filter", "", "Workflow executions are included in the returned results based on whether their start times are within the range specified by this filter.")
		swf_listOpenWorkflowExecutionsCmd.Flags().String("tag-filter", "", "If specified, only executions that have the matching tag are listed.")
		swf_listOpenWorkflowExecutionsCmd.Flags().String("type-filter", "", "If specified, only executions of the type specified in the filter are returned.")
		swf_listOpenWorkflowExecutionsCmd.MarkFlagRequired("domain")
		swf_listOpenWorkflowExecutionsCmd.MarkFlagRequired("start-time-filter")
	})
	swfCmd.AddCommand(swf_listOpenWorkflowExecutionsCmd)
}
