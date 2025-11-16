package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_listVpcEndpointsForDomainCmd = &cobra.Command{
	Use:   "list-vpc-endpoints-for-domain",
	Short: "Retrieves all Amazon OpenSearch Service-managed VPC endpoints associated with a particular domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_listVpcEndpointsForDomainCmd).Standalone()

	opensearch_listVpcEndpointsForDomainCmd.Flags().String("domain-name", "", "The name of the domain to list associated VPC endpoints for.")
	opensearch_listVpcEndpointsForDomainCmd.Flags().String("next-token", "", "If your initial `ListEndpointsForDomain` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListEndpointsForDomain` operations, which returns results in the next page.")
	opensearch_listVpcEndpointsForDomainCmd.MarkFlagRequired("domain-name")
	opensearchCmd.AddCommand(opensearch_listVpcEndpointsForDomainCmd)
}
