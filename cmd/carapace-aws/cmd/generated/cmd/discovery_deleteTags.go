package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_deleteTagsCmd = &cobra.Command{
	Use:   "delete-tags",
	Short: "Deletes the association between configuration items and one or more tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_deleteTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_deleteTagsCmd).Standalone()

		discovery_deleteTagsCmd.Flags().String("configuration-ids", "", "A list of configuration items with tags that you want to delete.")
		discovery_deleteTagsCmd.Flags().String("tags", "", "Tags that you want to delete from one or more configuration items.")
		discovery_deleteTagsCmd.MarkFlagRequired("configuration-ids")
	})
	discoveryCmd.AddCommand(discovery_deleteTagsCmd)
}
