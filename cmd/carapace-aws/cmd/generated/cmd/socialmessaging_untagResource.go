package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(socialmessaging_untagResourceCmd).Standalone()

		socialmessaging_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to remove tags from.")
		socialmessaging_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to remove from the resource.")
		socialmessaging_untagResourceCmd.MarkFlagRequired("resource-arn")
		socialmessaging_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	socialmessagingCmd.AddCommand(socialmessaging_untagResourceCmd)
}
