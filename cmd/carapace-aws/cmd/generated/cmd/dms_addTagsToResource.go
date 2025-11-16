package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_addTagsToResourceCmd = &cobra.Command{
	Use:   "add-tags-to-resource",
	Short: "Adds metadata tags to an DMS resource, including replication instance, endpoint, subnet group, and migration task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_addTagsToResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_addTagsToResourceCmd).Standalone()

		dms_addTagsToResourceCmd.Flags().String("resource-arn", "", "Identifies the DMS resource to which tags should be added.")
		dms_addTagsToResourceCmd.Flags().String("tags", "", "One or more tags to be assigned to the resource.")
		dms_addTagsToResourceCmd.MarkFlagRequired("resource-arn")
		dms_addTagsToResourceCmd.MarkFlagRequired("tags")
	})
	dmsCmd.AddCommand(dms_addTagsToResourceCmd)
}
