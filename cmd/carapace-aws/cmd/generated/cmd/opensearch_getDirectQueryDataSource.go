package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_getDirectQueryDataSourceCmd = &cobra.Command{
	Use:   "get-direct-query-data-source",
	Short: "Returns detailed configuration information for a specific direct query data source in Amazon OpenSearch Service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_getDirectQueryDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_getDirectQueryDataSourceCmd).Standalone()

		opensearch_getDirectQueryDataSourceCmd.Flags().String("data-source-name", "", "A unique, user-defined label that identifies the data source within your OpenSearch Service environment.")
		opensearch_getDirectQueryDataSourceCmd.MarkFlagRequired("data-source-name")
	})
	opensearchCmd.AddCommand(opensearch_getDirectQueryDataSourceCmd)
}
