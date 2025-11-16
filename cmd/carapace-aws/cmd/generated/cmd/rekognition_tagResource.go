package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more key-value tags to an Amazon Rekognition collection, stream processor, or Custom Labels model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_tagResourceCmd).Standalone()

	rekognition_tagResourceCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the model, collection, or stream processor that you want to assign the tags to.")
	rekognition_tagResourceCmd.Flags().String("tags", "", "The key-value tags to assign to the resource.")
	rekognition_tagResourceCmd.MarkFlagRequired("resource-arn")
	rekognition_tagResourceCmd.MarkFlagRequired("tags")
	rekognitionCmd.AddCommand(rekognition_tagResourceCmd)
}
