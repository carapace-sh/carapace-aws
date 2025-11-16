package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVpcEndpointServiceConfigurationCmd = &cobra.Command{
	Use:   "modify-vpc-endpoint-service-configuration",
	Short: "Modifies the attributes of the specified VPC endpoint service configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVpcEndpointServiceConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyVpcEndpointServiceConfigurationCmd).Standalone()

		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().Bool("acceptance-required", false, "Indicates whether requests to create an endpoint to the service must be accepted.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().String("add-gateway-load-balancer-arns", "", "The Amazon Resource Names (ARNs) of Gateway Load Balancers to add to the service configuration.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().String("add-network-load-balancer-arns", "", "The Amazon Resource Names (ARNs) of Network Load Balancers to add to the service configuration.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().String("add-supported-ip-address-types", "", "The IP address types to add to the service configuration.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().String("add-supported-regions", "", "The supported Regions to add to the service configuration.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().Bool("no-acceptance-required", false, "Indicates whether requests to create an endpoint to the service must be accepted.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().Bool("no-remove-private-dns-name", false, "(Interface endpoint configuration) Removes the private DNS name of the endpoint service.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().String("private-dns-name", "", "(Interface endpoint configuration) The private DNS name to assign to the endpoint service.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().String("remove-gateway-load-balancer-arns", "", "The Amazon Resource Names (ARNs) of Gateway Load Balancers to remove from the service configuration.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().String("remove-network-load-balancer-arns", "", "The Amazon Resource Names (ARNs) of Network Load Balancers to remove from the service configuration.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().Bool("remove-private-dns-name", false, "(Interface endpoint configuration) Removes the private DNS name of the endpoint service.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().String("remove-supported-ip-address-types", "", "The IP address types to remove from the service configuration.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().String("remove-supported-regions", "", "The supported Regions to remove from the service configuration.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flags().String("service-id", "", "The ID of the service.")
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flag("no-acceptance-required").Hidden = true
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyVpcEndpointServiceConfigurationCmd.Flag("no-remove-private-dns-name").Hidden = true
		ec2_modifyVpcEndpointServiceConfigurationCmd.MarkFlagRequired("service-id")
	})
	ec2Cmd.AddCommand(ec2_modifyVpcEndpointServiceConfigurationCmd)
}
