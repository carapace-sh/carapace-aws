package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_listVpcEndpointsCmd = &cobra.Command{
	Use:   "list-vpc-endpoints",
	Short: "Retrieves all Amazon OpenSearch Service-managed VPC endpoints in the current Amazon Web Services account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_listVpcEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_listVpcEndpointsCmd).Standalone()

		opensearch_listVpcEndpointsCmd.Flags().String("next-token", "", "If your initial `ListVpcEndpoints` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListVpcEndpoints` operations, which returns results in the next page.")
	})
	opensearchCmd.AddCommand(opensearch_listVpcEndpointsCmd)
}
