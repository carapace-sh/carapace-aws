package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_removeTagsCmd = &cobra.Command{
	Use:   "remove-tags",
	Short: "Removes the specified set of tags from an Amazon OpenSearch Service domain, data source, or application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_removeTagsCmd).Standalone()

	opensearch_removeTagsCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the domain, data source, or application from which you want to delete the specified tags.")
	opensearch_removeTagsCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the domain, data source, or application.")
	opensearch_removeTagsCmd.MarkFlagRequired("arn")
	opensearch_removeTagsCmd.MarkFlagRequired("tag-keys")
	opensearchCmd.AddCommand(opensearch_removeTagsCmd)
}
