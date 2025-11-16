package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_createTagsCmd = &cobra.Command{
	Use:   "create-tags",
	Short: "Creates one or more tags for configuration items.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_createTagsCmd).Standalone()

	discovery_createTagsCmd.Flags().String("configuration-ids", "", "A list of configuration items that you want to tag.")
	discovery_createTagsCmd.Flags().String("tags", "", "Tags that you want to associate with one or more configuration items.")
	discovery_createTagsCmd.MarkFlagRequired("configuration-ids")
	discovery_createTagsCmd.MarkFlagRequired("tags")
	discoveryCmd.AddCommand(discovery_createTagsCmd)
}
