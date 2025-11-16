package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more key-value tags to a Kinesis Analytics application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_tagResourceCmd).Standalone()

	kinesisanalytics_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the application to assign the tags.")
	kinesisanalytics_tagResourceCmd.Flags().String("tags", "", "The key-value tags to assign to the application.")
	kinesisanalytics_tagResourceCmd.MarkFlagRequired("resource-arn")
	kinesisanalytics_tagResourceCmd.MarkFlagRequired("tags")
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_tagResourceCmd)
}
