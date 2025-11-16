package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_listTagsCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "Returns all resource tags for an Amazon OpenSearch Service domain, data source, or application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_listTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_listTagsCmd).Standalone()

		opensearch_listTagsCmd.Flags().String("arn", "", "Amazon Resource Name (ARN) for the domain, data source, or application to view tags for.")
		opensearch_listTagsCmd.MarkFlagRequired("arn")
	})
	opensearchCmd.AddCommand(opensearch_listTagsCmd)
}
