package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more key-value tags to a Managed Service for Apache Flink application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_tagResourceCmd).Standalone()

		kinesisanalyticsv2_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the application to assign the tags.")
		kinesisanalyticsv2_tagResourceCmd.Flags().String("tags", "", "The key-value tags to assign to the application.")
		kinesisanalyticsv2_tagResourceCmd.MarkFlagRequired("resource-arn")
		kinesisanalyticsv2_tagResourceCmd.MarkFlagRequired("tags")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_tagResourceCmd)
}
