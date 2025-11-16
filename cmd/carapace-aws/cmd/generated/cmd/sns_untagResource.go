package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove tags from the specified Amazon SNS topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_untagResourceCmd).Standalone()

		sns_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the topic from which to remove tags.")
		sns_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the specified topic.")
		sns_untagResourceCmd.MarkFlagRequired("resource-arn")
		sns_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	snsCmd.AddCommand(sns_untagResourceCmd)
}
