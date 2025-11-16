package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_createVpcEndpointCmd = &cobra.Command{
	Use:   "create-vpc-endpoint",
	Short: "Creates an Amazon OpenSearch Service-managed VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_createVpcEndpointCmd).Standalone()

	opensearch_createVpcEndpointCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
	opensearch_createVpcEndpointCmd.Flags().String("domain-arn", "", "The Amazon Resource Name (ARN) of the domain to create the endpoint for.")
	opensearch_createVpcEndpointCmd.Flags().String("vpc-options", "", "Options to specify the subnets and security groups for the endpoint.")
	opensearch_createVpcEndpointCmd.MarkFlagRequired("domain-arn")
	opensearch_createVpcEndpointCmd.MarkFlagRequired("vpc-options")
	opensearchCmd.AddCommand(opensearch_createVpcEndpointCmd)
}
