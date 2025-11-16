package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the given tags (metadata) from the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_untagResourceCmd).Standalone()

	iotanalytics_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource whose tags you want to remove.")
	iotanalytics_untagResourceCmd.Flags().String("tag-keys", "", "The keys of those tags which you want to remove.")
	iotanalytics_untagResourceCmd.MarkFlagRequired("resource-arn")
	iotanalytics_untagResourceCmd.MarkFlagRequired("tag-keys")
	iotanalyticsCmd.AddCommand(iotanalytics_untagResourceCmd)
}
