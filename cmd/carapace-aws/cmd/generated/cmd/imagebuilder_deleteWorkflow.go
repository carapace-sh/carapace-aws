package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_deleteWorkflowCmd = &cobra.Command{
	Use:   "delete-workflow",
	Short: "Deletes a specific workflow resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_deleteWorkflowCmd).Standalone()

	imagebuilder_deleteWorkflowCmd.Flags().String("workflow-build-version-arn", "", "The Amazon Resource Name (ARN) of the workflow resource to delete.")
	imagebuilder_deleteWorkflowCmd.MarkFlagRequired("workflow-build-version-arn")
	imagebuilderCmd.AddCommand(imagebuilder_deleteWorkflowCmd)
}
