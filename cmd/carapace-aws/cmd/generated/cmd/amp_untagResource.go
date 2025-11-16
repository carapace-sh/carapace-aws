package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from an Amazon Managed Service for Prometheus resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_untagResourceCmd).Standalone()

		amp_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource from which to remove a tag.")
		amp_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to remove.")
		amp_untagResourceCmd.MarkFlagRequired("resource-arn")
		amp_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	ampCmd.AddCommand(amp_untagResourceCmd)
}
