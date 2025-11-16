package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_updateDomainConfigCmd = &cobra.Command{
	Use:   "update-domain-config",
	Short: "Modifies the cluster configuration of the specified Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_updateDomainConfigCmd).Standalone()

	opensearch_updateDomainConfigCmd.Flags().String("access-policies", "", "Identity and Access Management (IAM) access policy as a JSON-formatted string.")
	opensearch_updateDomainConfigCmd.Flags().String("advanced-options", "", "Key-value pairs to specify advanced configuration options.")
	opensearch_updateDomainConfigCmd.Flags().String("advanced-security-options", "", "Options for fine-grained access control.")
	opensearch_updateDomainConfigCmd.Flags().String("aimloptions", "", "Options for all machine learning features for the specified domain.")
	opensearch_updateDomainConfigCmd.Flags().String("auto-tune-options", "", "Options for Auto-Tune.")
	opensearch_updateDomainConfigCmd.Flags().String("cluster-config", "", "Changes that you want to make to the cluster configuration, such as the instance type and number of EC2 instances.")
	opensearch_updateDomainConfigCmd.Flags().String("cognito-options", "", "Key-value pairs to configure Amazon Cognito authentication for OpenSearch Dashboards.")
	opensearch_updateDomainConfigCmd.Flags().String("domain-endpoint-options", "", "Additional options for the domain endpoint, such as whether to require HTTPS for all traffic.")
	opensearch_updateDomainConfigCmd.Flags().String("domain-name", "", "The name of the domain that you're updating.")
	opensearch_updateDomainConfigCmd.Flags().String("dry-run", "", "This flag, when set to True, specifies whether the `UpdateDomain` request should return the results of a dry run analysis without actually applying the change.")
	opensearch_updateDomainConfigCmd.Flags().String("dry-run-mode", "", "The type of dry run to perform.")
	opensearch_updateDomainConfigCmd.Flags().String("ebsoptions", "", "The type and size of the EBS volume to attach to instances in the domain.")
	opensearch_updateDomainConfigCmd.Flags().String("encryption-at-rest-options", "", "Encryption at rest options for the domain.")
	opensearch_updateDomainConfigCmd.Flags().String("identity-center-options", "", "")
	opensearch_updateDomainConfigCmd.Flags().String("ipaddress-type", "", "Specify either dual stack or IPv4 as your IP address type.")
	opensearch_updateDomainConfigCmd.Flags().String("log-publishing-options", "", "Options to publish OpenSearch logs to Amazon CloudWatch Logs.")
	opensearch_updateDomainConfigCmd.Flags().String("node-to-node-encryption-options", "", "Node-to-node encryption options for the domain.")
	opensearch_updateDomainConfigCmd.Flags().String("off-peak-window-options", "", "Off-peak window options for the domain.")
	opensearch_updateDomainConfigCmd.Flags().String("snapshot-options", "", "Option to set the time, in UTC format, for the daily automated snapshot.")
	opensearch_updateDomainConfigCmd.Flags().String("software-update-options", "", "Service software update options for the domain.")
	opensearch_updateDomainConfigCmd.Flags().String("vpcoptions", "", "Options to specify the subnets and security groups for a VPC endpoint.")
	opensearch_updateDomainConfigCmd.MarkFlagRequired("domain-name")
	opensearchCmd.AddCommand(opensearch_updateDomainConfigCmd)
}
