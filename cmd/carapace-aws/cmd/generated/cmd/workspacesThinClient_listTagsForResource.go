package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_listTagsForResourceCmd).Standalone()

	workspacesThinClient_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which you want to retrieve tags.")
	workspacesThinClient_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	workspacesThinClientCmd.AddCommand(workspacesThinClient_listTagsForResourceCmd)
}
