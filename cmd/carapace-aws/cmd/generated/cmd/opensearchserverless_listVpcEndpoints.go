package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_listVpcEndpointsCmd = &cobra.Command{
	Use:   "list-vpc-endpoints",
	Short: "Returns the OpenSearch Serverless-managed interface VPC endpoints associated with the current account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_listVpcEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_listVpcEndpointsCmd).Standalone()

		opensearchserverless_listVpcEndpointsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		opensearchserverless_listVpcEndpointsCmd.Flags().String("next-token", "", "If your initial `ListVpcEndpoints` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListVpcEndpoints` operations, which returns results in the next page.")
		opensearchserverless_listVpcEndpointsCmd.Flags().String("vpc-endpoint-filters", "", "Filter the results according to the current status of the VPC endpoint.")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_listVpcEndpointsCmd)
}
