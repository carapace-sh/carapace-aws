package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_createVpcEndpointCmd = &cobra.Command{
	Use:   "create-vpc-endpoint",
	Short: "Creates an OpenSearch Serverless-managed interface VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_createVpcEndpointCmd).Standalone()

	opensearchserverless_createVpcEndpointCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
	opensearchserverless_createVpcEndpointCmd.Flags().String("name", "", "The name of the interface endpoint.")
	opensearchserverless_createVpcEndpointCmd.Flags().String("security-group-ids", "", "The unique identifiers of the security groups that define the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.")
	opensearchserverless_createVpcEndpointCmd.Flags().String("subnet-ids", "", "The ID of one or more subnets from which you'll access OpenSearch Serverless.")
	opensearchserverless_createVpcEndpointCmd.Flags().String("vpc-id", "", "The ID of the VPC from which you'll access OpenSearch Serverless.")
	opensearchserverless_createVpcEndpointCmd.MarkFlagRequired("name")
	opensearchserverless_createVpcEndpointCmd.MarkFlagRequired("subnet-ids")
	opensearchserverless_createVpcEndpointCmd.MarkFlagRequired("vpc-id")
	opensearchserverlessCmd.AddCommand(opensearchserverless_createVpcEndpointCmd)
}
