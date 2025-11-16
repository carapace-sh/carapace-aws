package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_listDataSourcesCmd = &cobra.Command{
	Use:   "list-data-sources",
	Short: "Lists direct-query data sources for a specific domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_listDataSourcesCmd).Standalone()

	opensearch_listDataSourcesCmd.Flags().String("domain-name", "", "The name of the domain.")
	opensearch_listDataSourcesCmd.MarkFlagRequired("domain-name")
	opensearchCmd.AddCommand(opensearch_listDataSourcesCmd)
}
