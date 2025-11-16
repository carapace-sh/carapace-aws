package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves the list of key-value tags assigned to the application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_listTagsForResourceCmd).Standalone()

	kinesisanalyticsv2_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the application for which to retrieve tags.")
	kinesisanalyticsv2_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_listTagsForResourceCmd)
}
