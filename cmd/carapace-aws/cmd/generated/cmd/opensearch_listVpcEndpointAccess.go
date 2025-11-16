package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_listVpcEndpointAccessCmd = &cobra.Command{
	Use:   "list-vpc-endpoint-access",
	Short: "Retrieves information about each Amazon Web Services principal that is allowed to access a given Amazon OpenSearch Service domain through the use of an interface VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_listVpcEndpointAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_listVpcEndpointAccessCmd).Standalone()

		opensearch_listVpcEndpointAccessCmd.Flags().String("domain-name", "", "The name of the OpenSearch Service domain to retrieve access information for.")
		opensearch_listVpcEndpointAccessCmd.Flags().String("next-token", "", "If your initial `ListVpcEndpointAccess` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListVpcEndpointAccess` operations, which returns results in the next page.")
		opensearch_listVpcEndpointAccessCmd.MarkFlagRequired("domain-name")
	})
	opensearchCmd.AddCommand(opensearch_listVpcEndpointAccessCmd)
}
