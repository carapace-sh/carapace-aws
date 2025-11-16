package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_updateDirectQueryDataSourceCmd = &cobra.Command{
	Use:   "update-direct-query-data-source",
	Short: "Updates the configuration or properties of an existing direct query data source in Amazon OpenSearch Service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_updateDirectQueryDataSourceCmd).Standalone()

	opensearch_updateDirectQueryDataSourceCmd.Flags().String("data-source-name", "", "A unique, user-defined label to identify the data source within your OpenSearch Service environment.")
	opensearch_updateDirectQueryDataSourceCmd.Flags().String("data-source-type", "", "The supported Amazon Web Services service that you want to use as the source for direct queries in OpenSearch Service.")
	opensearch_updateDirectQueryDataSourceCmd.Flags().String("description", "", "An optional text field for providing additional context and details about the data source.")
	opensearch_updateDirectQueryDataSourceCmd.Flags().String("open-search-arns", "", "A list of Amazon Resource Names (ARNs) for the OpenSearch collections that are associated with the direct query data source.")
	opensearch_updateDirectQueryDataSourceCmd.MarkFlagRequired("data-source-name")
	opensearch_updateDirectQueryDataSourceCmd.MarkFlagRequired("data-source-type")
	opensearch_updateDirectQueryDataSourceCmd.MarkFlagRequired("open-search-arns")
	opensearchCmd.AddCommand(opensearch_updateDirectQueryDataSourceCmd)
}
