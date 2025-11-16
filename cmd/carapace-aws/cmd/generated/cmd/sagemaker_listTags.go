package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listTagsCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "Returns the tags for the specified SageMaker resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listTagsCmd).Standalone()

		sagemaker_listTagsCmd.Flags().String("max-results", "", "Maximum number of tags to return.")
		sagemaker_listTagsCmd.Flags().String("next-token", "", "If the response to the previous `ListTags` request is truncated, SageMaker returns this token.")
		sagemaker_listTagsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource whose tags you want to retrieve.")
		sagemaker_listTagsCmd.MarkFlagRequired("resource-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_listTagsCmd)
}
