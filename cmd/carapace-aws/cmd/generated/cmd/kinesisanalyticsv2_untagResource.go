package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from a Managed Service for Apache Flink application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_untagResourceCmd).Standalone()

		kinesisanalyticsv2_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the Managed Service for Apache Flink application from which to remove the tags.")
		kinesisanalyticsv2_untagResourceCmd.Flags().String("tag-keys", "", "A list of keys of tags to remove from the specified application.")
		kinesisanalyticsv2_untagResourceCmd.MarkFlagRequired("resource-arn")
		kinesisanalyticsv2_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_untagResourceCmd)
}
