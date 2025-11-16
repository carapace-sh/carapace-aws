package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_tagResourceCmd).Standalone()

	workspacesThinClient_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to tag.")
	workspacesThinClient_tagResourceCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign to the resource.")
	workspacesThinClient_tagResourceCmd.MarkFlagRequired("resource-arn")
	workspacesThinClient_tagResourceCmd.MarkFlagRequired("tags")
	workspacesThinClientCmd.AddCommand(workspacesThinClient_tagResourceCmd)
}
