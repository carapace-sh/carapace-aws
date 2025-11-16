package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_registerWorkflowTypeCmd = &cobra.Command{
	Use:   "register-workflow-type",
	Short: "Registers a new *workflow type* and its configuration settings in the specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_registerWorkflowTypeCmd).Standalone()

	swf_registerWorkflowTypeCmd.Flags().String("default-child-policy", "", "If set, specifies the default policy to use for the child workflow executions when a workflow execution of this type is terminated, by calling the [TerminateWorkflowExecution]() action explicitly or due to an expired timeout.")
	swf_registerWorkflowTypeCmd.Flags().String("default-execution-start-to-close-timeout", "", "If set, specifies the default maximum duration for executions of this workflow type.")
	swf_registerWorkflowTypeCmd.Flags().String("default-lambda-role", "", "The default IAM role attached to this workflow type.")
	swf_registerWorkflowTypeCmd.Flags().String("default-task-list", "", "If set, specifies the default task list to use for scheduling decision tasks for executions of this workflow type.")
	swf_registerWorkflowTypeCmd.Flags().String("default-task-priority", "", "The default task priority to assign to the workflow type.")
	swf_registerWorkflowTypeCmd.Flags().String("default-task-start-to-close-timeout", "", "If set, specifies the default maximum duration of decision tasks for this workflow type.")
	swf_registerWorkflowTypeCmd.Flags().String("description", "", "Textual description of the workflow type.")
	swf_registerWorkflowTypeCmd.Flags().String("domain", "", "The name of the domain in which to register the workflow type.")
	swf_registerWorkflowTypeCmd.Flags().String("name", "", "The name of the workflow type.")
	swf_registerWorkflowTypeCmd.Flags().String("version", "", "The version of the workflow type.")
	swf_registerWorkflowTypeCmd.MarkFlagRequired("domain")
	swf_registerWorkflowTypeCmd.MarkFlagRequired("name")
	swf_registerWorkflowTypeCmd.MarkFlagRequired("version")
	swfCmd.AddCommand(swf_registerWorkflowTypeCmd)
}
