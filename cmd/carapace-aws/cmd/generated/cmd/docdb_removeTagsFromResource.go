package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_removeTagsFromResourceCmd = &cobra.Command{
	Use:   "remove-tags-from-resource",
	Short: "Removes metadata tags from an Amazon DocumentDB resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_removeTagsFromResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_removeTagsFromResourceCmd).Standalone()

		docdb_removeTagsFromResourceCmd.Flags().String("resource-name", "", "The Amazon DocumentDB resource that the tags are removed from.")
		docdb_removeTagsFromResourceCmd.Flags().String("tag-keys", "", "The tag key (name) of the tag to be removed.")
		docdb_removeTagsFromResourceCmd.MarkFlagRequired("resource-name")
		docdb_removeTagsFromResourceCmd.MarkFlagRequired("tag-keys")
	})
	docdbCmd.AddCommand(docdb_removeTagsFromResourceCmd)
}
