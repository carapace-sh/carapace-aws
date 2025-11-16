package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_updateVpcEndpointCmd = &cobra.Command{
	Use:   "update-vpc-endpoint",
	Short: "Updates an OpenSearch Serverless-managed interface endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_updateVpcEndpointCmd).Standalone()

	opensearchserverless_updateVpcEndpointCmd.Flags().String("add-security-group-ids", "", "The unique identifiers of the security groups to add to the endpoint.")
	opensearchserverless_updateVpcEndpointCmd.Flags().String("add-subnet-ids", "", "The ID of one or more subnets to add to the endpoint.")
	opensearchserverless_updateVpcEndpointCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
	opensearchserverless_updateVpcEndpointCmd.Flags().String("id", "", "The unique identifier of the interface endpoint to update.")
	opensearchserverless_updateVpcEndpointCmd.Flags().String("remove-security-group-ids", "", "The unique identifiers of the security groups to remove from the endpoint.")
	opensearchserverless_updateVpcEndpointCmd.Flags().String("remove-subnet-ids", "", "The unique identifiers of the subnets to remove from the endpoint.")
	opensearchserverless_updateVpcEndpointCmd.MarkFlagRequired("id")
	opensearchserverlessCmd.AddCommand(opensearchserverless_updateVpcEndpointCmd)
}
