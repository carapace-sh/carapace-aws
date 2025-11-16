package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_describeElasticsearchDomainsCmd = &cobra.Command{
	Use:   "describe-elasticsearch-domains",
	Short: "Returns domain configuration information about the specified Elasticsearch domains, including the domain ID, domain endpoint, and domain ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_describeElasticsearchDomainsCmd).Standalone()

	es_describeElasticsearchDomainsCmd.Flags().String("domain-names", "", "The Elasticsearch domains for which you want information.")
	es_describeElasticsearchDomainsCmd.MarkFlagRequired("domain-names")
	esCmd.AddCommand(es_describeElasticsearchDomainsCmd)
}
