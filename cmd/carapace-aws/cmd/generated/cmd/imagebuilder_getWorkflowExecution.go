package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getWorkflowExecutionCmd = &cobra.Command{
	Use:   "get-workflow-execution",
	Short: "Get the runtime information that was logged for a specific runtime instance of the workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getWorkflowExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_getWorkflowExecutionCmd).Standalone()

		imagebuilder_getWorkflowExecutionCmd.Flags().String("workflow-execution-id", "", "Use the unique identifier for a runtime instance of the workflow to get runtime details.")
		imagebuilder_getWorkflowExecutionCmd.MarkFlagRequired("workflow-execution-id")
	})
	imagebuilderCmd.AddCommand(imagebuilder_getWorkflowExecutionCmd)
}
