package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove one or more tags (keys and values) from a specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_untagResourceCmd).Standalone()

		applicationInsights_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the application that you want to remove one or more tags from.")
		applicationInsights_untagResourceCmd.Flags().String("tag-keys", "", "The tags (tag keys) that you want to remove from the resource.")
		applicationInsights_untagResourceCmd.MarkFlagRequired("resource-arn")
		applicationInsights_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_untagResourceCmd)
}
