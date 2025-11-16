package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_deleteDirectQueryDataSourceCmd = &cobra.Command{
	Use:   "delete-direct-query-data-source",
	Short: "Deletes a previously configured direct query data source from Amazon OpenSearch Service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_deleteDirectQueryDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_deleteDirectQueryDataSourceCmd).Standalone()

		opensearch_deleteDirectQueryDataSourceCmd.Flags().String("data-source-name", "", "A unique, user-defined label to identify the data source within your OpenSearch Service environment.")
		opensearch_deleteDirectQueryDataSourceCmd.MarkFlagRequired("data-source-name")
	})
	opensearchCmd.AddCommand(opensearch_deleteDirectQueryDataSourceCmd)
}
