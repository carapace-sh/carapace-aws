package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from an Amazon Rekognition collection, stream processor, or Custom Labels model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_untagResourceCmd).Standalone()

	rekognition_untagResourceCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the model, collection, or stream processor that you want to remove the tags from.")
	rekognition_untagResourceCmd.Flags().String("tag-keys", "", "A list of the tags that you want to remove.")
	rekognition_untagResourceCmd.MarkFlagRequired("resource-arn")
	rekognition_untagResourceCmd.MarkFlagRequired("tag-keys")
	rekognitionCmd.AddCommand(rekognition_untagResourceCmd)
}
