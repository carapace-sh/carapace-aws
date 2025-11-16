package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_describeWorkflowExecutionCmd = &cobra.Command{
	Use:   "describe-workflow-execution",
	Short: "Returns information about the specified workflow execution including its type and some statistics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_describeWorkflowExecutionCmd).Standalone()

	swf_describeWorkflowExecutionCmd.Flags().String("domain", "", "The name of the domain containing the workflow execution.")
	swf_describeWorkflowExecutionCmd.Flags().String("execution", "", "The workflow execution to describe.")
	swf_describeWorkflowExecutionCmd.MarkFlagRequired("domain")
	swf_describeWorkflowExecutionCmd.MarkFlagRequired("execution")
	swfCmd.AddCommand(swf_describeWorkflowExecutionCmd)
}
