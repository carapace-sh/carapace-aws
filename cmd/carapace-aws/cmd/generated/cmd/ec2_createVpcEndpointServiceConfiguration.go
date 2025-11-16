package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createVpcEndpointServiceConfigurationCmd = &cobra.Command{
	Use:   "create-vpc-endpoint-service-configuration",
	Short: "Creates a VPC endpoint service to which service consumers (Amazon Web Services accounts, users, and IAM roles) can connect.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createVpcEndpointServiceConfigurationCmd).Standalone()

	ec2_createVpcEndpointServiceConfigurationCmd.Flags().Bool("acceptance-required", false, "Indicates whether requests from service consumers to create an endpoint to your service must be accepted manually.")
	ec2_createVpcEndpointServiceConfigurationCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createVpcEndpointServiceConfigurationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVpcEndpointServiceConfigurationCmd.Flags().String("gateway-load-balancer-arns", "", "The Amazon Resource Names (ARNs) of the Gateway Load Balancers.")
	ec2_createVpcEndpointServiceConfigurationCmd.Flags().String("network-load-balancer-arns", "", "The Amazon Resource Names (ARNs) of the Network Load Balancers.")
	ec2_createVpcEndpointServiceConfigurationCmd.Flags().Bool("no-acceptance-required", false, "Indicates whether requests from service consumers to create an endpoint to your service must be accepted manually.")
	ec2_createVpcEndpointServiceConfigurationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVpcEndpointServiceConfigurationCmd.Flags().String("private-dns-name", "", "(Interface endpoint configuration) The private DNS name to assign to the VPC endpoint service.")
	ec2_createVpcEndpointServiceConfigurationCmd.Flags().String("supported-ip-address-types", "", "The supported IP address types.")
	ec2_createVpcEndpointServiceConfigurationCmd.Flags().String("supported-regions", "", "The Regions from which service consumers can access the service.")
	ec2_createVpcEndpointServiceConfigurationCmd.Flags().String("tag-specifications", "", "The tags to associate with the service.")
	ec2_createVpcEndpointServiceConfigurationCmd.Flag("no-acceptance-required").Hidden = true
	ec2_createVpcEndpointServiceConfigurationCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createVpcEndpointServiceConfigurationCmd)
}
