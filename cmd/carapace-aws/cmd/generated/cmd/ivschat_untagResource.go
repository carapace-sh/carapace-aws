package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from the resource with the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivschat_untagResourceCmd).Standalone()

		ivschat_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to be untagged.")
		ivschat_untagResourceCmd.Flags().String("tag-keys", "", "Array of tags to be removed.")
		ivschat_untagResourceCmd.MarkFlagRequired("resource-arn")
		ivschat_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	ivschatCmd.AddCommand(ivschat_untagResourceCmd)
}
