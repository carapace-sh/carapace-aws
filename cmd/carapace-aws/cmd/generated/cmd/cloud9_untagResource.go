package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud9_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from an Cloud9 development environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud9_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloud9_untagResourceCmd).Standalone()

		cloud9_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Cloud9 development environment to remove tags from.")
		cloud9_untagResourceCmd.Flags().String("tag-keys", "", "The tag names of the tags to remove from the given Cloud9 development environment.")
		cloud9_untagResourceCmd.MarkFlagRequired("resource-arn")
		cloud9_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	cloud9Cmd.AddCommand(cloud9_untagResourceCmd)
}
