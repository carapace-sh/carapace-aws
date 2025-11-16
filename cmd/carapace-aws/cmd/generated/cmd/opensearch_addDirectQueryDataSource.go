package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_addDirectQueryDataSourceCmd = &cobra.Command{
	Use:   "add-direct-query-data-source",
	Short: "Adds a new data source in Amazon OpenSearch Service so that you can perform direct queries on external data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_addDirectQueryDataSourceCmd).Standalone()

	opensearch_addDirectQueryDataSourceCmd.Flags().String("data-source-name", "", "A unique, user-defined label to identify the data source within your OpenSearch Service environment.")
	opensearch_addDirectQueryDataSourceCmd.Flags().String("data-source-type", "", "The supported Amazon Web Services service that you want to use as the source for direct queries in OpenSearch Service.")
	opensearch_addDirectQueryDataSourceCmd.Flags().String("description", "", "An optional text field for providing additional context and details about the data source.")
	opensearch_addDirectQueryDataSourceCmd.Flags().String("open-search-arns", "", "A list of Amazon Resource Names (ARNs) for the OpenSearch collections that are associated with the direct query data source.")
	opensearch_addDirectQueryDataSourceCmd.Flags().String("tag-list", "", "")
	opensearch_addDirectQueryDataSourceCmd.MarkFlagRequired("data-source-name")
	opensearch_addDirectQueryDataSourceCmd.MarkFlagRequired("data-source-type")
	opensearch_addDirectQueryDataSourceCmd.MarkFlagRequired("open-search-arns")
	opensearchCmd.AddCommand(opensearch_addDirectQueryDataSourceCmd)
}
