package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from a Kinesis Analytics application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_untagResourceCmd).Standalone()

	kinesisanalytics_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the Kinesis Analytics application from which to remove the tags.")
	kinesisanalytics_untagResourceCmd.Flags().String("tag-keys", "", "A list of keys of tags to remove from the specified application.")
	kinesisanalytics_untagResourceCmd.MarkFlagRequired("resource-arn")
	kinesisanalytics_untagResourceCmd.MarkFlagRequired("tag-keys")
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_untagResourceCmd)
}
