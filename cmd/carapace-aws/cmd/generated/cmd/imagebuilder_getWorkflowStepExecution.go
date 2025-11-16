package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getWorkflowStepExecutionCmd = &cobra.Command{
	Use:   "get-workflow-step-execution",
	Short: "Get the runtime information that was logged for a specific runtime instance of the workflow step.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getWorkflowStepExecutionCmd).Standalone()

	imagebuilder_getWorkflowStepExecutionCmd.Flags().String("step-execution-id", "", "Use the unique identifier for a specific runtime instance of the workflow step to get runtime details for that step.")
	imagebuilder_getWorkflowStepExecutionCmd.MarkFlagRequired("step-execution-id")
	imagebuilderCmd.AddCommand(imagebuilder_getWorkflowStepExecutionCmd)
}
