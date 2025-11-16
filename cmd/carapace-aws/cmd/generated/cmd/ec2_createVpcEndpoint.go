package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createVpcEndpointCmd = &cobra.Command{
	Use:   "create-vpc-endpoint",
	Short: "Creates a VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createVpcEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createVpcEndpointCmd).Standalone()

		ec2_createVpcEndpointCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createVpcEndpointCmd.Flags().String("dns-options", "", "The DNS options for the endpoint.")
		ec2_createVpcEndpointCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createVpcEndpointCmd.Flags().String("ip-address-type", "", "The IP address type for the endpoint.")
		ec2_createVpcEndpointCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createVpcEndpointCmd.Flags().Bool("no-private-dns-enabled", false, "(Interface endpoint) Indicates whether to associate a private hosted zone with the specified VPC.")
		ec2_createVpcEndpointCmd.Flags().String("policy-document", "", "(Interface and gateway endpoints) A policy to attach to the endpoint that controls access to the service.")
		ec2_createVpcEndpointCmd.Flags().Bool("private-dns-enabled", false, "(Interface endpoint) Indicates whether to associate a private hosted zone with the specified VPC.")
		ec2_createVpcEndpointCmd.Flags().String("resource-configuration-arn", "", "The Amazon Resource Name (ARN) of a resource configuration that will be associated with the VPC endpoint of type resource.")
		ec2_createVpcEndpointCmd.Flags().String("route-table-ids", "", "(Gateway endpoint) The route table IDs.")
		ec2_createVpcEndpointCmd.Flags().String("security-group-ids", "", "(Interface endpoint) The IDs of the security groups to associate with the endpoint network interfaces.")
		ec2_createVpcEndpointCmd.Flags().String("service-name", "", "The name of the endpoint service.")
		ec2_createVpcEndpointCmd.Flags().String("service-network-arn", "", "The Amazon Resource Name (ARN) of a service network that will be associated with the VPC endpoint of type service-network.")
		ec2_createVpcEndpointCmd.Flags().String("service-region", "", "The Region where the service is hosted.")
		ec2_createVpcEndpointCmd.Flags().String("subnet-configurations", "", "The subnet configurations for the endpoint.")
		ec2_createVpcEndpointCmd.Flags().String("subnet-ids", "", "(Interface and Gateway Load Balancer endpoints) The IDs of the subnets in which to create endpoint network interfaces.")
		ec2_createVpcEndpointCmd.Flags().String("tag-specifications", "", "The tags to associate with the endpoint.")
		ec2_createVpcEndpointCmd.Flags().String("vpc-endpoint-type", "", "The type of endpoint.")
		ec2_createVpcEndpointCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
		ec2_createVpcEndpointCmd.Flag("no-dry-run").Hidden = true
		ec2_createVpcEndpointCmd.Flag("no-private-dns-enabled").Hidden = true
		ec2_createVpcEndpointCmd.MarkFlagRequired("vpc-id")
	})
	ec2Cmd.AddCommand(ec2_createVpcEndpointCmd)
}
