package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags in an Amazon Rekognition collection, stream processor, or Custom Labels model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_listTagsForResourceCmd).Standalone()

		rekognition_listTagsForResourceCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the model, collection, or stream processor that contains the tags that you want a list of.")
		rekognition_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	rekognitionCmd.AddCommand(rekognition_listTagsForResourceCmd)
}
