package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_requestCancelWorkflowExecutionCmd = &cobra.Command{
	Use:   "request-cancel-workflow-execution",
	Short: "Records a `WorkflowExecutionCancelRequested` event in the currently running workflow execution identified by the given domain, workflowId, and runId.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_requestCancelWorkflowExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_requestCancelWorkflowExecutionCmd).Standalone()

		swf_requestCancelWorkflowExecutionCmd.Flags().String("domain", "", "The name of the domain containing the workflow execution to cancel.")
		swf_requestCancelWorkflowExecutionCmd.Flags().String("run-id", "", "The runId of the workflow execution to cancel.")
		swf_requestCancelWorkflowExecutionCmd.Flags().String("workflow-id", "", "The workflowId of the workflow execution to cancel.")
		swf_requestCancelWorkflowExecutionCmd.MarkFlagRequired("domain")
		swf_requestCancelWorkflowExecutionCmd.MarkFlagRequired("workflow-id")
	})
	swfCmd.AddCommand(swf_requestCancelWorkflowExecutionCmd)
}
