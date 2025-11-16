package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listWorkflowExecutionsCmd = &cobra.Command{
	Use:   "list-workflow-executions",
	Short: "Returns a list of workflow runtime instance metadata objects for a specific image build version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listWorkflowExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_listWorkflowExecutionsCmd).Standalone()

		imagebuilder_listWorkflowExecutionsCmd.Flags().String("image-build-version-arn", "", "List all workflow runtime instances for the specified image build version resource ARN.")
		imagebuilder_listWorkflowExecutionsCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
		imagebuilder_listWorkflowExecutionsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
		imagebuilder_listWorkflowExecutionsCmd.MarkFlagRequired("image-build-version-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_listWorkflowExecutionsCmd)
}
