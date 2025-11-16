package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_untagResourceCmd).Standalone()

		controltower_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		controltower_untagResourceCmd.Flags().String("tag-keys", "", "Tag keys to be removed from the resource.")
		controltower_untagResourceCmd.MarkFlagRequired("resource-arn")
		controltower_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	controltowerCmd.AddCommand(controltower_untagResourceCmd)
}
