package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_updateVpcEndpointCmd = &cobra.Command{
	Use:   "update-vpc-endpoint",
	Short: "Modifies an Amazon OpenSearch Service-managed interface VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_updateVpcEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_updateVpcEndpointCmd).Standalone()

		opensearch_updateVpcEndpointCmd.Flags().String("vpc-endpoint-id", "", "The unique identifier of the endpoint.")
		opensearch_updateVpcEndpointCmd.Flags().String("vpc-options", "", "The security groups and/or subnets to add, remove, or modify.")
		opensearch_updateVpcEndpointCmd.MarkFlagRequired("vpc-endpoint-id")
		opensearch_updateVpcEndpointCmd.MarkFlagRequired("vpc-options")
	})
	opensearchCmd.AddCommand(opensearch_updateVpcEndpointCmd)
}
