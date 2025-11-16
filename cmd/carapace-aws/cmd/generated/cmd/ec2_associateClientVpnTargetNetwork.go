package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateClientVpnTargetNetworkCmd = &cobra.Command{
	Use:   "associate-client-vpn-target-network",
	Short: "Associates a target network with a Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateClientVpnTargetNetworkCmd).Standalone()

	ec2_associateClientVpnTargetNetworkCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_associateClientVpnTargetNetworkCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint.")
	ec2_associateClientVpnTargetNetworkCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateClientVpnTargetNetworkCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateClientVpnTargetNetworkCmd.Flags().String("subnet-id", "", "The ID of the subnet to associate with the Client VPN endpoint.")
	ec2_associateClientVpnTargetNetworkCmd.MarkFlagRequired("client-vpn-endpoint-id")
	ec2_associateClientVpnTargetNetworkCmd.Flag("no-dry-run").Hidden = true
	ec2_associateClientVpnTargetNetworkCmd.MarkFlagRequired("subnet-id")
	ec2Cmd.AddCommand(ec2_associateClientVpnTargetNetworkCmd)
}
