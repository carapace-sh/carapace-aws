package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_listDirectQueryDataSourcesCmd = &cobra.Command{
	Use:   "list-direct-query-data-sources",
	Short: "Lists an inventory of all the direct query data sources that you have configured within Amazon OpenSearch Service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_listDirectQueryDataSourcesCmd).Standalone()

	opensearch_listDirectQueryDataSourcesCmd.Flags().String("next-token", "", "")
	opensearchCmd.AddCommand(opensearch_listDirectQueryDataSourcesCmd)
}
