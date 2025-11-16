package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves the list of key-value tags assigned to the application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalytics_listTagsForResourceCmd).Standalone()

		kinesisanalytics_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the application for which to retrieve tags.")
		kinesisanalytics_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_listTagsForResourceCmd)
}
