package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_describeElasticsearchDomainCmd = &cobra.Command{
	Use:   "describe-elasticsearch-domain",
	Short: "Returns domain configuration information about the specified Elasticsearch domain, including the domain ID, domain endpoint, and domain ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_describeElasticsearchDomainCmd).Standalone()

	es_describeElasticsearchDomainCmd.Flags().String("domain-name", "", "The name of the Elasticsearch domain for which you want information.")
	es_describeElasticsearchDomainCmd.MarkFlagRequired("domain-name")
	esCmd.AddCommand(es_describeElasticsearchDomainCmd)
}
