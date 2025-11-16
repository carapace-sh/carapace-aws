package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all metadata tags attached to an DMS resource, including replication instance, endpoint, subnet group, and migration task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_listTagsForResourceCmd).Standalone()

	dms_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) string that uniquely identifies the DMS resource to list tags for.")
	dms_listTagsForResourceCmd.Flags().String("resource-arn-list", "", "List of ARNs that identify multiple DMS resources that you want to list tags for.")
	dmsCmd.AddCommand(dms_listTagsForResourceCmd)
}
