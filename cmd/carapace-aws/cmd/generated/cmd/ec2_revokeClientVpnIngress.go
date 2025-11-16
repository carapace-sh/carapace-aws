package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_revokeClientVpnIngressCmd = &cobra.Command{
	Use:   "revoke-client-vpn-ingress",
	Short: "Removes an ingress authorization rule from a Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_revokeClientVpnIngressCmd).Standalone()

	ec2_revokeClientVpnIngressCmd.Flags().String("access-group-id", "", "The ID of the Active Directory group for which to revoke access.")
	ec2_revokeClientVpnIngressCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint with which the authorization rule is associated.")
	ec2_revokeClientVpnIngressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_revokeClientVpnIngressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_revokeClientVpnIngressCmd.Flags().Bool("no-revoke-all-groups", false, "Indicates whether access should be revoked for all groups for a single `TargetNetworkCidr` that earlier authorized ingress for all groups using `AuthorizeAllGroups`.")
	ec2_revokeClientVpnIngressCmd.Flags().Bool("revoke-all-groups", false, "Indicates whether access should be revoked for all groups for a single `TargetNetworkCidr` that earlier authorized ingress for all groups using `AuthorizeAllGroups`.")
	ec2_revokeClientVpnIngressCmd.Flags().String("target-network-cidr", "", "The IPv4 address range, in CIDR notation, of the network for which access is being removed.")
	ec2_revokeClientVpnIngressCmd.MarkFlagRequired("client-vpn-endpoint-id")
	ec2_revokeClientVpnIngressCmd.Flag("no-dry-run").Hidden = true
	ec2_revokeClientVpnIngressCmd.Flag("no-revoke-all-groups").Hidden = true
	ec2_revokeClientVpnIngressCmd.MarkFlagRequired("target-network-cidr")
	ec2Cmd.AddCommand(ec2_revokeClientVpnIngressCmd)
}
