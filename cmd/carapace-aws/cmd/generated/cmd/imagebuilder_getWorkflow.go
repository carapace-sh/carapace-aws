package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getWorkflowCmd = &cobra.Command{
	Use:   "get-workflow",
	Short: "Get a workflow resource object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getWorkflowCmd).Standalone()

	imagebuilder_getWorkflowCmd.Flags().String("workflow-build-version-arn", "", "The Amazon Resource Name (ARN) of the workflow resource that you want to get.")
	imagebuilder_getWorkflowCmd.MarkFlagRequired("workflow-build-version-arn")
	imagebuilderCmd.AddCommand(imagebuilder_getWorkflowCmd)
}
