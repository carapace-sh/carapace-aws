package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_listElasticsearchVersionsCmd = &cobra.Command{
	Use:   "list-elasticsearch-versions",
	Short: "List all supported Elasticsearch versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_listElasticsearchVersionsCmd).Standalone()

	es_listElasticsearchVersionsCmd.Flags().String("max-results", "", "Set this value to limit the number of results returned.")
	es_listElasticsearchVersionsCmd.Flags().String("next-token", "", "")
	esCmd.AddCommand(es_listElasticsearchVersionsCmd)
}
