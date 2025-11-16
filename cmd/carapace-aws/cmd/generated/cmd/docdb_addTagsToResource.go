package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_addTagsToResourceCmd = &cobra.Command{
	Use:   "add-tags-to-resource",
	Short: "Adds metadata tags to an Amazon DocumentDB resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_addTagsToResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_addTagsToResourceCmd).Standalone()

		docdb_addTagsToResourceCmd.Flags().String("resource-name", "", "The Amazon DocumentDB resource that the tags are added to.")
		docdb_addTagsToResourceCmd.Flags().String("tags", "", "The tags to be assigned to the Amazon DocumentDB resource.")
		docdb_addTagsToResourceCmd.MarkFlagRequired("resource-name")
		docdb_addTagsToResourceCmd.MarkFlagRequired("tags")
	})
	docdbCmd.AddCommand(docdb_addTagsToResourceCmd)
}
