package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_removeTagsFromResourceCmd = &cobra.Command{
	Use:   "remove-tags-from-resource",
	Short: "Removes metadata tags from an Amazon Neptune resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_removeTagsFromResourceCmd).Standalone()

	neptune_removeTagsFromResourceCmd.Flags().String("resource-name", "", "The Amazon Neptune resource that the tags are removed from.")
	neptune_removeTagsFromResourceCmd.Flags().String("tag-keys", "", "The tag key (name) of the tag to be removed.")
	neptune_removeTagsFromResourceCmd.MarkFlagRequired("resource-name")
	neptune_removeTagsFromResourceCmd.MarkFlagRequired("tag-keys")
	neptuneCmd.AddCommand(neptune_removeTagsFromResourceCmd)
}
