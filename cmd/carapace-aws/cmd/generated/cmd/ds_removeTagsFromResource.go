package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_removeTagsFromResourceCmd = &cobra.Command{
	Use:   "remove-tags-from-resource",
	Short: "Removes tags from a directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_removeTagsFromResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_removeTagsFromResourceCmd).Standalone()

		ds_removeTagsFromResourceCmd.Flags().String("resource-id", "", "Identifier (ID) of the directory from which to remove the tag.")
		ds_removeTagsFromResourceCmd.Flags().String("tag-keys", "", "The tag key (name) of the tag to be removed.")
		ds_removeTagsFromResourceCmd.MarkFlagRequired("resource-id")
		ds_removeTagsFromResourceCmd.MarkFlagRequired("tag-keys")
	})
	dsCmd.AddCommand(ds_removeTagsFromResourceCmd)
}
