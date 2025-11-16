package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_listElasticsearchInstanceTypesCmd = &cobra.Command{
	Use:   "list-elasticsearch-instance-types",
	Short: "List all Elasticsearch instance types that are supported for given ElasticsearchVersion",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_listElasticsearchInstanceTypesCmd).Standalone()

	es_listElasticsearchInstanceTypesCmd.Flags().String("domain-name", "", "DomainName represents the name of the Domain that we are trying to modify.")
	es_listElasticsearchInstanceTypesCmd.Flags().String("elasticsearch-version", "", "Version of Elasticsearch for which list of supported elasticsearch instance types are needed.")
	es_listElasticsearchInstanceTypesCmd.Flags().String("max-results", "", "Set this value to limit the number of results returned.")
	es_listElasticsearchInstanceTypesCmd.Flags().String("next-token", "", "NextToken should be sent in case if earlier API call produced result containing NextToken.")
	es_listElasticsearchInstanceTypesCmd.MarkFlagRequired("elasticsearch-version")
	esCmd.AddCommand(es_listElasticsearchInstanceTypesCmd)
}
