package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag or tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_untagResourceCmd).Standalone()

	workspacesThinClient_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to untag.")
	workspacesThinClient_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the key-value pairs for the tag or tags you want to remove from the specified resource.")
	workspacesThinClient_untagResourceCmd.MarkFlagRequired("resource-arn")
	workspacesThinClient_untagResourceCmd.MarkFlagRequired("tag-keys")
	workspacesThinClientCmd.AddCommand(workspacesThinClient_untagResourceCmd)
}
