package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_authorizeClientVpnIngressCmd = &cobra.Command{
	Use:   "authorize-client-vpn-ingress",
	Short: "Adds an ingress authorization rule to a Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_authorizeClientVpnIngressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_authorizeClientVpnIngressCmd).Standalone()

		ec2_authorizeClientVpnIngressCmd.Flags().String("access-group-id", "", "The ID of the group to grant access to, for example, the Active Directory group or identity provider (IdP) group.")
		ec2_authorizeClientVpnIngressCmd.Flags().Bool("authorize-all-groups", false, "Indicates whether to grant access to all clients.")
		ec2_authorizeClientVpnIngressCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_authorizeClientVpnIngressCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint.")
		ec2_authorizeClientVpnIngressCmd.Flags().String("description", "", "A brief description of the authorization rule.")
		ec2_authorizeClientVpnIngressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_authorizeClientVpnIngressCmd.Flags().Bool("no-authorize-all-groups", false, "Indicates whether to grant access to all clients.")
		ec2_authorizeClientVpnIngressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_authorizeClientVpnIngressCmd.Flags().String("target-network-cidr", "", "The IPv4 address range, in CIDR notation, of the network for which access is being authorized.")
		ec2_authorizeClientVpnIngressCmd.MarkFlagRequired("client-vpn-endpoint-id")
		ec2_authorizeClientVpnIngressCmd.Flag("no-authorize-all-groups").Hidden = true
		ec2_authorizeClientVpnIngressCmd.Flag("no-dry-run").Hidden = true
		ec2_authorizeClientVpnIngressCmd.MarkFlagRequired("target-network-cidr")
	})
	ec2Cmd.AddCommand(ec2_authorizeClientVpnIngressCmd)
}
