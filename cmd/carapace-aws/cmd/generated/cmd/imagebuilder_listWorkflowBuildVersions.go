package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listWorkflowBuildVersionsCmd = &cobra.Command{
	Use:   "list-workflow-build-versions",
	Short: "Returns a list of build versions for a specific workflow resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listWorkflowBuildVersionsCmd).Standalone()

	imagebuilder_listWorkflowBuildVersionsCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
	imagebuilder_listWorkflowBuildVersionsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	imagebuilder_listWorkflowBuildVersionsCmd.Flags().String("workflow-version-arn", "", "The Amazon Resource Name (ARN) of the workflow resource for which to get a list of build versions.")
	imagebuilderCmd.AddCommand(imagebuilder_listWorkflowBuildVersionsCmd)
}
