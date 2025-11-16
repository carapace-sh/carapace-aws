package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_startWorkflowExecutionCmd = &cobra.Command{
	Use:   "start-workflow-execution",
	Short: "Starts an execution of the workflow type in the specified domain using the provided `workflowId` and input data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_startWorkflowExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_startWorkflowExecutionCmd).Standalone()

		swf_startWorkflowExecutionCmd.Flags().String("child-policy", "", "If set, specifies the policy to use for the child workflow executions of this workflow execution if it is terminated, by calling the [TerminateWorkflowExecution]() action explicitly or due to an expired timeout.")
		swf_startWorkflowExecutionCmd.Flags().String("domain", "", "The name of the domain in which the workflow execution is created.")
		swf_startWorkflowExecutionCmd.Flags().String("execution-start-to-close-timeout", "", "The total duration for this workflow execution.")
		swf_startWorkflowExecutionCmd.Flags().String("input", "", "The input for the workflow execution.")
		swf_startWorkflowExecutionCmd.Flags().String("lambda-role", "", "The IAM role to attach to this workflow execution.")
		swf_startWorkflowExecutionCmd.Flags().String("tag-list", "", "The list of tags to associate with the workflow execution.")
		swf_startWorkflowExecutionCmd.Flags().String("task-list", "", "The task list to use for the decision tasks generated for this workflow execution.")
		swf_startWorkflowExecutionCmd.Flags().String("task-priority", "", "The task priority to use for this workflow execution.")
		swf_startWorkflowExecutionCmd.Flags().String("task-start-to-close-timeout", "", "Specifies the maximum duration of decision tasks for this workflow execution.")
		swf_startWorkflowExecutionCmd.Flags().String("workflow-id", "", "The user defined identifier associated with the workflow execution.")
		swf_startWorkflowExecutionCmd.Flags().String("workflow-type", "", "The type of the workflow to start.")
		swf_startWorkflowExecutionCmd.MarkFlagRequired("domain")
		swf_startWorkflowExecutionCmd.MarkFlagRequired("workflow-id")
		swf_startWorkflowExecutionCmd.MarkFlagRequired("workflow-type")
	})
	swfCmd.AddCommand(swf_startWorkflowExecutionCmd)
}
