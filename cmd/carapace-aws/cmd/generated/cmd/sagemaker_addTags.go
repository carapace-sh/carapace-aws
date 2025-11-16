package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_addTagsCmd = &cobra.Command{
	Use:   "add-tags",
	Short: "Adds or overwrites one or more tags for the specified SageMaker resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_addTagsCmd).Standalone()

	sagemaker_addTagsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to tag.")
	sagemaker_addTagsCmd.Flags().String("tags", "", "An array of key-value pairs.")
	sagemaker_addTagsCmd.MarkFlagRequired("resource-arn")
	sagemaker_addTagsCmd.MarkFlagRequired("tags")
	sagemakerCmd.AddCommand(sagemaker_addTagsCmd)
}
