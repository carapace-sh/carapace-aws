package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_listVpcEndpointAccessCmd = &cobra.Command{
	Use:   "list-vpc-endpoint-access",
	Short: "Retrieves information about each principal that is allowed to access a given Amazon OpenSearch Service domain through the use of an interface VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_listVpcEndpointAccessCmd).Standalone()

	es_listVpcEndpointAccessCmd.Flags().String("domain-name", "", "The name of the OpenSearch Service domain to retrieve access information for.")
	es_listVpcEndpointAccessCmd.Flags().String("next-token", "", "Provides an identifier to allow retrieval of paginated results.")
	es_listVpcEndpointAccessCmd.MarkFlagRequired("domain-name")
	esCmd.AddCommand(es_listVpcEndpointAccessCmd)
}
