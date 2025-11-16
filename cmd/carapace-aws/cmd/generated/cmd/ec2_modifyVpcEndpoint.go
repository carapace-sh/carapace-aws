package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVpcEndpointCmd = &cobra.Command{
	Use:   "modify-vpc-endpoint",
	Short: "Modifies attributes of a specified VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVpcEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyVpcEndpointCmd).Standalone()

		ec2_modifyVpcEndpointCmd.Flags().String("add-route-table-ids", "", "(Gateway endpoint) The IDs of the route tables to associate with the endpoint.")
		ec2_modifyVpcEndpointCmd.Flags().String("add-security-group-ids", "", "(Interface endpoint) The IDs of the security groups to associate with the endpoint network interfaces.")
		ec2_modifyVpcEndpointCmd.Flags().String("add-subnet-ids", "", "(Interface and Gateway Load Balancer endpoints) The IDs of the subnets in which to serve the endpoint.")
		ec2_modifyVpcEndpointCmd.Flags().String("dns-options", "", "The DNS options for the endpoint.")
		ec2_modifyVpcEndpointCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVpcEndpointCmd.Flags().String("ip-address-type", "", "The IP address type for the endpoint.")
		ec2_modifyVpcEndpointCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVpcEndpointCmd.Flags().Bool("no-private-dns-enabled", false, "(Interface endpoint) Indicates whether a private hosted zone is associated with the VPC.")
		ec2_modifyVpcEndpointCmd.Flags().Bool("no-reset-policy", false, "(Gateway endpoint) Specify `true` to reset the policy document to the default policy.")
		ec2_modifyVpcEndpointCmd.Flags().String("policy-document", "", "(Interface and gateway endpoints) A policy to attach to the endpoint that controls access to the service.")
		ec2_modifyVpcEndpointCmd.Flags().Bool("private-dns-enabled", false, "(Interface endpoint) Indicates whether a private hosted zone is associated with the VPC.")
		ec2_modifyVpcEndpointCmd.Flags().String("remove-route-table-ids", "", "(Gateway endpoint) The IDs of the route tables to disassociate from the endpoint.")
		ec2_modifyVpcEndpointCmd.Flags().String("remove-security-group-ids", "", "(Interface endpoint) The IDs of the security groups to disassociate from the endpoint network interfaces.")
		ec2_modifyVpcEndpointCmd.Flags().String("remove-subnet-ids", "", "(Interface endpoint) The IDs of the subnets from which to remove the endpoint.")
		ec2_modifyVpcEndpointCmd.Flags().Bool("reset-policy", false, "(Gateway endpoint) Specify `true` to reset the policy document to the default policy.")
		ec2_modifyVpcEndpointCmd.Flags().String("subnet-configurations", "", "The subnet configurations for the endpoint.")
		ec2_modifyVpcEndpointCmd.Flags().String("vpc-endpoint-id", "", "The ID of the endpoint.")
		ec2_modifyVpcEndpointCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyVpcEndpointCmd.Flag("no-private-dns-enabled").Hidden = true
		ec2_modifyVpcEndpointCmd.Flag("no-reset-policy").Hidden = true
		ec2_modifyVpcEndpointCmd.MarkFlagRequired("vpc-endpoint-id")
	})
	ec2Cmd.AddCommand(ec2_modifyVpcEndpointCmd)
}
