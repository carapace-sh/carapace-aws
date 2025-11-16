package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listWorkflowStepExecutionsCmd = &cobra.Command{
	Use:   "list-workflow-step-executions",
	Short: "Returns runtime data for each step in a runtime instance of the workflow that you specify in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listWorkflowStepExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_listWorkflowStepExecutionsCmd).Standalone()

		imagebuilder_listWorkflowStepExecutionsCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
		imagebuilder_listWorkflowStepExecutionsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
		imagebuilder_listWorkflowStepExecutionsCmd.Flags().String("workflow-execution-id", "", "The unique identifier that Image Builder assigned to keep track of runtime details when it ran the workflow.")
		imagebuilder_listWorkflowStepExecutionsCmd.MarkFlagRequired("workflow-execution-id")
	})
	imagebuilderCmd.AddCommand(imagebuilder_listWorkflowStepExecutionsCmd)
}
