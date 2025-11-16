package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_applySecurityGroupsToClientVpnTargetNetworkCmd = &cobra.Command{
	Use:   "apply-security-groups-to-client-vpn-target-network",
	Short: "Applies a security group to the association between the target network and the Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_applySecurityGroupsToClientVpnTargetNetworkCmd).Standalone()

	ec2_applySecurityGroupsToClientVpnTargetNetworkCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint.")
	ec2_applySecurityGroupsToClientVpnTargetNetworkCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_applySecurityGroupsToClientVpnTargetNetworkCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_applySecurityGroupsToClientVpnTargetNetworkCmd.Flags().String("security-group-ids", "", "The IDs of the security groups to apply to the associated target network.")
	ec2_applySecurityGroupsToClientVpnTargetNetworkCmd.Flags().String("vpc-id", "", "The ID of the VPC in which the associated target network is located.")
	ec2_applySecurityGroupsToClientVpnTargetNetworkCmd.MarkFlagRequired("client-vpn-endpoint-id")
	ec2_applySecurityGroupsToClientVpnTargetNetworkCmd.Flag("no-dry-run").Hidden = true
	ec2_applySecurityGroupsToClientVpnTargetNetworkCmd.MarkFlagRequired("security-group-ids")
	ec2_applySecurityGroupsToClientVpnTargetNetworkCmd.MarkFlagRequired("vpc-id")
	ec2Cmd.AddCommand(ec2_applySecurityGroupsToClientVpnTargetNetworkCmd)
}
