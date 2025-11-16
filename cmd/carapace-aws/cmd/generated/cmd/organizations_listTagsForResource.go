package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists tags that are attached to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_listTagsForResourceCmd).Standalone()

		organizations_listTagsForResourceCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
		organizations_listTagsForResourceCmd.Flags().String("resource-id", "", "The ID of the resource with the tags to list.")
		organizations_listTagsForResourceCmd.MarkFlagRequired("resource-id")
	})
	organizationsCmd.AddCommand(organizations_listTagsForResourceCmd)
}
