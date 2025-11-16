package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_removeTagsFromResourceCmd = &cobra.Command{
	Use:   "remove-tags-from-resource",
	Short: "Removes metadata tags from an DMS resource, including replication instance, endpoint, subnet group, and migration task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_removeTagsFromResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_removeTagsFromResourceCmd).Standalone()

		dms_removeTagsFromResourceCmd.Flags().String("resource-arn", "", "An DMS resource from which you want to remove tag(s).")
		dms_removeTagsFromResourceCmd.Flags().String("tag-keys", "", "The tag key (name) of the tag to be removed.")
		dms_removeTagsFromResourceCmd.MarkFlagRequired("resource-arn")
		dms_removeTagsFromResourceCmd.MarkFlagRequired("tag-keys")
	})
	dmsCmd.AddCommand(dms_removeTagsFromResourceCmd)
}
