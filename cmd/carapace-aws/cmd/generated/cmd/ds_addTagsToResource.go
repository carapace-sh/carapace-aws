package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_addTagsToResourceCmd = &cobra.Command{
	Use:   "add-tags-to-resource",
	Short: "Adds or overwrites one or more tags for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_addTagsToResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_addTagsToResourceCmd).Standalone()

		ds_addTagsToResourceCmd.Flags().String("resource-id", "", "Identifier (ID) for the directory to which to add the tag.")
		ds_addTagsToResourceCmd.Flags().String("tags", "", "The tags to be assigned to the directory.")
		ds_addTagsToResourceCmd.MarkFlagRequired("resource-id")
		ds_addTagsToResourceCmd.MarkFlagRequired("tags")
	})
	dsCmd.AddCommand(ds_addTagsToResourceCmd)
}
