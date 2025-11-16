package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove one or more tags (keys and values) from a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_untagResourceCmd).Standalone()

	mailmanager_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to untag.")
	mailmanager_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the key-value pairs for the tag or tags you want to remove from the specified resource.")
	mailmanager_untagResourceCmd.MarkFlagRequired("resource-arn")
	mailmanager_untagResourceCmd.MarkFlagRequired("tag-keys")
	mailmanagerCmd.AddCommand(mailmanager_untagResourceCmd)
}
