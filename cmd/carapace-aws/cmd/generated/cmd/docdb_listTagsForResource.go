package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags on an Amazon DocumentDB resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_listTagsForResourceCmd).Standalone()

		docdb_listTagsForResourceCmd.Flags().String("filters", "", "This parameter is not currently supported.")
		docdb_listTagsForResourceCmd.Flags().String("resource-name", "", "The Amazon DocumentDB resource with tags to be listed.")
		docdb_listTagsForResourceCmd.MarkFlagRequired("resource-name")
	})
	docdbCmd.AddCommand(docdb_listTagsForResourceCmd)
}
