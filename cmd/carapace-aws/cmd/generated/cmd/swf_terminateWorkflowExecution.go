package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_terminateWorkflowExecutionCmd = &cobra.Command{
	Use:   "terminate-workflow-execution",
	Short: "Records a `WorkflowExecutionTerminated` event and forces closure of the workflow execution identified by the given domain, runId, and workflowId.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_terminateWorkflowExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_terminateWorkflowExecutionCmd).Standalone()

		swf_terminateWorkflowExecutionCmd.Flags().String("child-policy", "", "If set, specifies the policy to use for the child workflow executions of the workflow execution being terminated.")
		swf_terminateWorkflowExecutionCmd.Flags().String("details", "", "Details for terminating the workflow execution.")
		swf_terminateWorkflowExecutionCmd.Flags().String("domain", "", "The domain of the workflow execution to terminate.")
		swf_terminateWorkflowExecutionCmd.Flags().String("reason", "", "A descriptive reason for terminating the workflow execution.")
		swf_terminateWorkflowExecutionCmd.Flags().String("run-id", "", "The runId of the workflow execution to terminate.")
		swf_terminateWorkflowExecutionCmd.Flags().String("workflow-id", "", "The workflowId of the workflow execution to terminate.")
		swf_terminateWorkflowExecutionCmd.MarkFlagRequired("domain")
		swf_terminateWorkflowExecutionCmd.MarkFlagRequired("workflow-id")
	})
	swfCmd.AddCommand(swf_terminateWorkflowExecutionCmd)
}
