package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_describeElasticsearchDomainConfigCmd = &cobra.Command{
	Use:   "describe-elasticsearch-domain-config",
	Short: "Provides cluster configuration information about the specified Elasticsearch domain, such as the state, creation date, update version, and update date for cluster options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_describeElasticsearchDomainConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_describeElasticsearchDomainConfigCmd).Standalone()

		es_describeElasticsearchDomainConfigCmd.Flags().String("domain-name", "", "The Elasticsearch domain that you want to get information about.")
		es_describeElasticsearchDomainConfigCmd.MarkFlagRequired("domain-name")
	})
	esCmd.AddCommand(es_describeElasticsearchDomainConfigCmd)
}
