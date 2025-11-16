package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_updateVpcEndpointCmd = &cobra.Command{
	Use:   "update-vpc-endpoint",
	Short: "Modifies an Amazon OpenSearch Service-managed interface VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_updateVpcEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_updateVpcEndpointCmd).Standalone()

		es_updateVpcEndpointCmd.Flags().String("vpc-endpoint-id", "", "Unique identifier of the VPC endpoint to be updated.")
		es_updateVpcEndpointCmd.Flags().String("vpc-options", "", "The security groups and/or subnets to add, remove, or modify.")
		es_updateVpcEndpointCmd.MarkFlagRequired("vpc-endpoint-id")
		es_updateVpcEndpointCmd.MarkFlagRequired("vpc-options")
	})
	esCmd.AddCommand(es_updateVpcEndpointCmd)
}
