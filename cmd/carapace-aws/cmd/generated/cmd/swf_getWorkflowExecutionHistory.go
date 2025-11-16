package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_getWorkflowExecutionHistoryCmd = &cobra.Command{
	Use:   "get-workflow-execution-history",
	Short: "Returns the history of the specified workflow execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_getWorkflowExecutionHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_getWorkflowExecutionHistoryCmd).Standalone()

		swf_getWorkflowExecutionHistoryCmd.Flags().String("domain", "", "The name of the domain containing the workflow execution.")
		swf_getWorkflowExecutionHistoryCmd.Flags().String("execution", "", "Specifies the workflow execution for which to return the history.")
		swf_getWorkflowExecutionHistoryCmd.Flags().String("maximum-page-size", "", "The maximum number of results that are returned per call.")
		swf_getWorkflowExecutionHistoryCmd.Flags().String("next-page-token", "", "If `NextPageToken` is returned there are more results available.")
		swf_getWorkflowExecutionHistoryCmd.Flags().String("reverse-order", "", "When set to `true`, returns the events in reverse order.")
		swf_getWorkflowExecutionHistoryCmd.MarkFlagRequired("domain")
		swf_getWorkflowExecutionHistoryCmd.MarkFlagRequired("execution")
	})
	swfCmd.AddCommand(swf_getWorkflowExecutionHistoryCmd)
}
