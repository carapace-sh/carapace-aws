package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_addTagsCmd = &cobra.Command{
	Use:   "add-tags",
	Short: "Attaches tags to an existing Amazon OpenSearch Service domain, data source, or application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_addTagsCmd).Standalone()

	opensearch_addTagsCmd.Flags().String("arn", "", "Amazon Resource Name (ARN) for the OpenSearch Service domain, data source, or application to which you want to attach resource tags.")
	opensearch_addTagsCmd.Flags().String("tag-list", "", "List of resource tags.")
	opensearch_addTagsCmd.MarkFlagRequired("arn")
	opensearch_addTagsCmd.MarkFlagRequired("tag-list")
	opensearchCmd.AddCommand(opensearch_addTagsCmd)
}
