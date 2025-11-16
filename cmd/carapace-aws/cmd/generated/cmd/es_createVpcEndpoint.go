package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_createVpcEndpointCmd = &cobra.Command{
	Use:   "create-vpc-endpoint",
	Short: "Creates an Amazon OpenSearch Service-managed VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_createVpcEndpointCmd).Standalone()

	es_createVpcEndpointCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
	es_createVpcEndpointCmd.Flags().String("domain-arn", "", "The Amazon Resource Name (ARN) of the domain to grant access to.")
	es_createVpcEndpointCmd.Flags().String("vpc-options", "", "Options to specify the subnets and security groups for the endpoint.")
	es_createVpcEndpointCmd.MarkFlagRequired("domain-arn")
	es_createVpcEndpointCmd.MarkFlagRequired("vpc-options")
	esCmd.AddCommand(es_createVpcEndpointCmd)
}
