package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags on a directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_listTagsForResourceCmd).Standalone()

	ds_listTagsForResourceCmd.Flags().String("limit", "", "Reserved for future use.")
	ds_listTagsForResourceCmd.Flags().String("next-token", "", "Reserved for future use.")
	ds_listTagsForResourceCmd.Flags().String("resource-id", "", "Identifier (ID) of the directory for which you want to retrieve tags.")
	ds_listTagsForResourceCmd.MarkFlagRequired("resource-id")
	dsCmd.AddCommand(ds_listTagsForResourceCmd)
}
