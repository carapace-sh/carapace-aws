package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from a resource using the resource's ARN and tag to remove.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_untagResourceCmd).Standalone()

		deadline_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to remove the tag from.")
		deadline_untagResourceCmd.Flags().String("tag-keys", "", "They keys of the tag.")
		deadline_untagResourceCmd.MarkFlagRequired("resource-arn")
		deadline_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	deadlineCmd.AddCommand(deadline_untagResourceCmd)
}
