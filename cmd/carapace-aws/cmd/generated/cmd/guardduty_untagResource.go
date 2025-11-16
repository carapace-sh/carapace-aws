package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_untagResourceCmd).Standalone()

		guardduty_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the resource to remove tags from.")
		guardduty_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys to remove from the resource.")
		guardduty_untagResourceCmd.MarkFlagRequired("resource-arn")
		guardduty_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	guarddutyCmd.AddCommand(guardduty_untagResourceCmd)
}
