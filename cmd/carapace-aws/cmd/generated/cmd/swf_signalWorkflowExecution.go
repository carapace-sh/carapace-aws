package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_signalWorkflowExecutionCmd = &cobra.Command{
	Use:   "signal-workflow-execution",
	Short: "Records a `WorkflowExecutionSignaled` event in the workflow execution history and creates a decision task for the workflow execution identified by the given domain, workflowId and runId.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_signalWorkflowExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_signalWorkflowExecutionCmd).Standalone()

		swf_signalWorkflowExecutionCmd.Flags().String("domain", "", "The name of the domain containing the workflow execution to signal.")
		swf_signalWorkflowExecutionCmd.Flags().String("input", "", "Data to attach to the `WorkflowExecutionSignaled` event in the target workflow execution's history.")
		swf_signalWorkflowExecutionCmd.Flags().String("run-id", "", "The runId of the workflow execution to signal.")
		swf_signalWorkflowExecutionCmd.Flags().String("signal-name", "", "The name of the signal.")
		swf_signalWorkflowExecutionCmd.Flags().String("workflow-id", "", "The workflowId of the workflow execution to signal.")
		swf_signalWorkflowExecutionCmd.MarkFlagRequired("domain")
		swf_signalWorkflowExecutionCmd.MarkFlagRequired("signal-name")
		swf_signalWorkflowExecutionCmd.MarkFlagRequired("workflow-id")
	})
	swfCmd.AddCommand(swf_signalWorkflowExecutionCmd)
}
