package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_revokeVpcEndpointAccessCmd = &cobra.Command{
	Use:   "revoke-vpc-endpoint-access",
	Short: "Revokes access to an Amazon OpenSearch Service domain that was provided through an interface VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_revokeVpcEndpointAccessCmd).Standalone()

	opensearch_revokeVpcEndpointAccessCmd.Flags().String("account", "", "The account ID to revoke access from.")
	opensearch_revokeVpcEndpointAccessCmd.Flags().String("domain-name", "", "The name of the OpenSearch Service domain.")
	opensearch_revokeVpcEndpointAccessCmd.Flags().String("service", "", "The service SP to revoke access from.")
	opensearch_revokeVpcEndpointAccessCmd.MarkFlagRequired("domain-name")
	opensearchCmd.AddCommand(opensearch_revokeVpcEndpointAccessCmd)
}
