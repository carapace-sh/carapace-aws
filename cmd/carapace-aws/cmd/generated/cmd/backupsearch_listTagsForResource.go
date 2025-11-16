package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupsearch_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "This operation returns the tags for a resource type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupsearch_listTagsForResourceCmd).Standalone()

	backupsearch_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the resource.&gt;")
	backupsearch_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	backupsearchCmd.AddCommand(backupsearch_listTagsForResourceCmd)
}
