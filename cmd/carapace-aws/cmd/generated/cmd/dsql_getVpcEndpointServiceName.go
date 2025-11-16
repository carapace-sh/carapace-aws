package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsql_getVpcEndpointServiceNameCmd = &cobra.Command{
	Use:   "get-vpc-endpoint-service-name",
	Short: "Retrieves the VPC endpoint service name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsql_getVpcEndpointServiceNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsql_getVpcEndpointServiceNameCmd).Standalone()

		dsql_getVpcEndpointServiceNameCmd.Flags().String("identifier", "", "The ID of the cluster to retrieve.")
		dsql_getVpcEndpointServiceNameCmd.MarkFlagRequired("identifier")
	})
	dsqlCmd.AddCommand(dsql_getVpcEndpointServiceNameCmd)
}
