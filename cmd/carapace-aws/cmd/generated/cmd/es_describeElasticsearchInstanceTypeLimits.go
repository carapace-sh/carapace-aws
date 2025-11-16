package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_describeElasticsearchInstanceTypeLimitsCmd = &cobra.Command{
	Use:   "describe-elasticsearch-instance-type-limits",
	Short: "Describe Elasticsearch Limits for a given InstanceType and ElasticsearchVersion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_describeElasticsearchInstanceTypeLimitsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_describeElasticsearchInstanceTypeLimitsCmd).Standalone()

		es_describeElasticsearchInstanceTypeLimitsCmd.Flags().String("domain-name", "", "DomainName represents the name of the Domain that we are trying to modify.")
		es_describeElasticsearchInstanceTypeLimitsCmd.Flags().String("elasticsearch-version", "", "Version of Elasticsearch for which `Limits` are needed.")
		es_describeElasticsearchInstanceTypeLimitsCmd.Flags().String("instance-type", "", "The instance type for an Elasticsearch cluster for which Elasticsearch `Limits` are needed.")
		es_describeElasticsearchInstanceTypeLimitsCmd.MarkFlagRequired("elasticsearch-version")
		es_describeElasticsearchInstanceTypeLimitsCmd.MarkFlagRequired("instance-type")
	})
	esCmd.AddCommand(es_describeElasticsearchInstanceTypeLimitsCmd)
}
