package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_listVpcEndpointsForDomainCmd = &cobra.Command{
	Use:   "list-vpc-endpoints-for-domain",
	Short: "Retrieves all Amazon OpenSearch Service-managed VPC endpoints associated with a particular domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_listVpcEndpointsForDomainCmd).Standalone()

	es_listVpcEndpointsForDomainCmd.Flags().String("domain-name", "", "Name of the ElasticSearch domain whose VPC endpoints are to be listed.")
	es_listVpcEndpointsForDomainCmd.Flags().String("next-token", "", "Provides an identifier to allow retrieval of paginated results.")
	es_listVpcEndpointsForDomainCmd.MarkFlagRequired("domain-name")
	esCmd.AddCommand(es_listVpcEndpointsForDomainCmd)
}
