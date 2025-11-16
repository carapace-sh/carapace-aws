package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rbin_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Unassigns a tag from a retention rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rbin_untagResourceCmd).Standalone()

	rbin_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the retention rule.")
	rbin_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys of the tags to unassign.")
	rbin_untagResourceCmd.MarkFlagRequired("resource-arn")
	rbin_untagResourceCmd.MarkFlagRequired("tag-keys")
	rbinCmd.AddCommand(rbin_untagResourceCmd)
}
