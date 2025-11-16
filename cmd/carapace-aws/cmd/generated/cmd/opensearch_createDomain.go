package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_createDomainCmd = &cobra.Command{
	Use:   "create-domain",
	Short: "Creates an Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_createDomainCmd).Standalone()

	opensearch_createDomainCmd.Flags().String("access-policies", "", "Identity and Access Management (IAM) policy document specifying the access policies for the new domain.")
	opensearch_createDomainCmd.Flags().String("advanced-options", "", "Key-value pairs to specify advanced configuration options.")
	opensearch_createDomainCmd.Flags().String("advanced-security-options", "", "Options for fine-grained access control.")
	opensearch_createDomainCmd.Flags().String("aimloptions", "", "Options for all machine learning features for the specified domain.")
	opensearch_createDomainCmd.Flags().String("auto-tune-options", "", "Options for Auto-Tune.")
	opensearch_createDomainCmd.Flags().String("cluster-config", "", "Container for the cluster configuration of a domain.")
	opensearch_createDomainCmd.Flags().String("cognito-options", "", "Key-value pairs to configure Amazon Cognito authentication.")
	opensearch_createDomainCmd.Flags().String("domain-endpoint-options", "", "Additional options for the domain endpoint, such as whether to require HTTPS for all traffic.")
	opensearch_createDomainCmd.Flags().String("domain-name", "", "Name of the OpenSearch Service domain to create.")
	opensearch_createDomainCmd.Flags().String("ebsoptions", "", "Container for the parameters required to enable EBS-based storage for an OpenSearch Service domain.")
	opensearch_createDomainCmd.Flags().String("encryption-at-rest-options", "", "Key-value pairs to enable encryption at rest.")
	opensearch_createDomainCmd.Flags().String("engine-version", "", "String of format Elasticsearch\\_X.Y or OpenSearch\\_X.Y to specify the engine version for the OpenSearch Service domain.")
	opensearch_createDomainCmd.Flags().String("identity-center-options", "", "Configuration options for enabling and managing IAM Identity Center integration within a domain.")
	opensearch_createDomainCmd.Flags().String("ipaddress-type", "", "Specify either dual stack or IPv4 as your IP address type.")
	opensearch_createDomainCmd.Flags().String("log-publishing-options", "", "Key-value pairs to configure log publishing.")
	opensearch_createDomainCmd.Flags().String("node-to-node-encryption-options", "", "Enables node-to-node encryption.")
	opensearch_createDomainCmd.Flags().String("off-peak-window-options", "", "Specifies a daily 10-hour time block during which OpenSearch Service can perform configuration changes on the domain, including service software updates and Auto-Tune enhancements that require a blue/green deployment.")
	opensearch_createDomainCmd.Flags().String("snapshot-options", "", "DEPRECATED.")
	opensearch_createDomainCmd.Flags().String("software-update-options", "", "Software update options for the domain.")
	opensearch_createDomainCmd.Flags().String("tag-list", "", "List of tags to add to the domain upon creation.")
	opensearch_createDomainCmd.Flags().String("vpcoptions", "", "Container for the values required to configure VPC access domains.")
	opensearch_createDomainCmd.MarkFlagRequired("domain-name")
	opensearchCmd.AddCommand(opensearch_createDomainCmd)
}
