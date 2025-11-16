package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteTagsCmd = &cobra.Command{
	Use:   "delete-tags",
	Short: "Deletes the specified tags from an SageMaker resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteTagsCmd).Standalone()

		sagemaker_deleteTagsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource whose tags you want to delete.")
		sagemaker_deleteTagsCmd.Flags().String("tag-keys", "", "An array or one or more tag keys to delete.")
		sagemaker_deleteTagsCmd.MarkFlagRequired("resource-arn")
		sagemaker_deleteTagsCmd.MarkFlagRequired("tag-keys")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteTagsCmd)
}
