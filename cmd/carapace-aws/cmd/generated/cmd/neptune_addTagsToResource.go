package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_addTagsToResourceCmd = &cobra.Command{
	Use:   "add-tags-to-resource",
	Short: "Adds metadata tags to an Amazon Neptune resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_addTagsToResourceCmd).Standalone()

	neptune_addTagsToResourceCmd.Flags().String("resource-name", "", "The Amazon Neptune resource that the tags are added to.")
	neptune_addTagsToResourceCmd.Flags().String("tags", "", "The tags to be assigned to the Amazon Neptune resource.")
	neptune_addTagsToResourceCmd.MarkFlagRequired("resource-name")
	neptune_addTagsToResourceCmd.MarkFlagRequired("tags")
	neptuneCmd.AddCommand(neptune_addTagsToResourceCmd)
}
