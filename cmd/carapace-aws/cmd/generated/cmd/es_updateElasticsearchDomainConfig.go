package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_updateElasticsearchDomainConfigCmd = &cobra.Command{
	Use:   "update-elasticsearch-domain-config",
	Short: "Modifies the cluster configuration of the specified Elasticsearch domain, setting as setting the instance type and the number of instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_updateElasticsearchDomainConfigCmd).Standalone()

	es_updateElasticsearchDomainConfigCmd.Flags().String("access-policies", "", "IAM access policy as a JSON-formatted string.")
	es_updateElasticsearchDomainConfigCmd.Flags().String("advanced-options", "", "Modifies the advanced option to allow references to indices in an HTTP request body.")
	es_updateElasticsearchDomainConfigCmd.Flags().String("advanced-security-options", "", "Specifies advanced security options.")
	es_updateElasticsearchDomainConfigCmd.Flags().String("auto-tune-options", "", "Specifies Auto-Tune options.")
	es_updateElasticsearchDomainConfigCmd.Flags().String("cognito-options", "", "Options to specify the Cognito user and identity pools for Kibana authentication.")
	es_updateElasticsearchDomainConfigCmd.Flags().String("domain-endpoint-options", "", "Options to specify configuration that will be applied to the domain endpoint.")
	es_updateElasticsearchDomainConfigCmd.Flags().String("domain-name", "", "The name of the Elasticsearch domain that you are updating.")
	es_updateElasticsearchDomainConfigCmd.Flags().String("dry-run", "", "This flag, when set to True, specifies whether the `UpdateElasticsearchDomain` request should return the results of validation checks without actually applying the change.")
	es_updateElasticsearchDomainConfigCmd.Flags().String("ebsoptions", "", "Specify the type and size of the EBS volume that you want to use.")
	es_updateElasticsearchDomainConfigCmd.Flags().String("elasticsearch-cluster-config", "", "The type and number of instances to instantiate for the domain cluster.")
	es_updateElasticsearchDomainConfigCmd.Flags().String("encryption-at-rest-options", "", "Specifies the Encryption At Rest Options.")
	es_updateElasticsearchDomainConfigCmd.Flags().String("log-publishing-options", "", "Map of `LogType` and `LogPublishingOption`, each containing options to publish a given type of Elasticsearch log.")
	es_updateElasticsearchDomainConfigCmd.Flags().String("node-to-node-encryption-options", "", "Specifies the NodeToNodeEncryptionOptions.")
	es_updateElasticsearchDomainConfigCmd.Flags().String("snapshot-options", "", "Option to set the time, in UTC format, for the daily automated snapshot.")
	es_updateElasticsearchDomainConfigCmd.Flags().String("vpcoptions", "", "Options to specify the subnets and security groups for VPC endpoint.")
	es_updateElasticsearchDomainConfigCmd.MarkFlagRequired("domain-name")
	esCmd.AddCommand(es_updateElasticsearchDomainConfigCmd)
}
