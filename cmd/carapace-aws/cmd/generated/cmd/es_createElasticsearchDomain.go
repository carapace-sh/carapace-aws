package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_createElasticsearchDomainCmd = &cobra.Command{
	Use:   "create-elasticsearch-domain",
	Short: "Creates a new Elasticsearch domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_createElasticsearchDomainCmd).Standalone()

	es_createElasticsearchDomainCmd.Flags().String("access-policies", "", "IAM access policy as a JSON-formatted string.")
	es_createElasticsearchDomainCmd.Flags().String("advanced-options", "", "Option to allow references to indices in an HTTP request body.")
	es_createElasticsearchDomainCmd.Flags().String("advanced-security-options", "", "Specifies advanced security options.")
	es_createElasticsearchDomainCmd.Flags().String("auto-tune-options", "", "Specifies Auto-Tune options.")
	es_createElasticsearchDomainCmd.Flags().String("cognito-options", "", "Options to specify the Cognito user and identity pools for Kibana authentication.")
	es_createElasticsearchDomainCmd.Flags().String("domain-endpoint-options", "", "Options to specify configuration that will be applied to the domain endpoint.")
	es_createElasticsearchDomainCmd.Flags().String("domain-name", "", "The name of the Elasticsearch domain that you are creating.")
	es_createElasticsearchDomainCmd.Flags().String("ebsoptions", "", "Options to enable, disable and specify the type and size of EBS storage volumes.")
	es_createElasticsearchDomainCmd.Flags().String("elasticsearch-cluster-config", "", "Configuration options for an Elasticsearch domain.")
	es_createElasticsearchDomainCmd.Flags().String("elasticsearch-version", "", "String of format X.Y to specify version for the Elasticsearch domain eg.")
	es_createElasticsearchDomainCmd.Flags().String("encryption-at-rest-options", "", "Specifies the Encryption At Rest Options.")
	es_createElasticsearchDomainCmd.Flags().String("log-publishing-options", "", "Map of `LogType` and `LogPublishingOption`, each containing options to publish a given type of Elasticsearch log.")
	es_createElasticsearchDomainCmd.Flags().String("node-to-node-encryption-options", "", "Specifies the NodeToNodeEncryptionOptions.")
	es_createElasticsearchDomainCmd.Flags().String("snapshot-options", "", "Option to set time, in UTC format, of the daily automated snapshot.")
	es_createElasticsearchDomainCmd.Flags().String("tag-list", "", "A list of `Tag` added during domain creation.")
	es_createElasticsearchDomainCmd.Flags().String("vpcoptions", "", "Options to specify the subnets and security groups for VPC endpoint.")
	es_createElasticsearchDomainCmd.MarkFlagRequired("domain-name")
	esCmd.AddCommand(es_createElasticsearchDomainCmd)
}
