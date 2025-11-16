package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateClientVpnTargetNetworkCmd = &cobra.Command{
	Use:   "disassociate-client-vpn-target-network",
	Short: "Disassociates a target network from the specified Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateClientVpnTargetNetworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disassociateClientVpnTargetNetworkCmd).Standalone()

		ec2_disassociateClientVpnTargetNetworkCmd.Flags().String("association-id", "", "The ID of the target network association.")
		ec2_disassociateClientVpnTargetNetworkCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint from which to disassociate the target network.")
		ec2_disassociateClientVpnTargetNetworkCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateClientVpnTargetNetworkCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateClientVpnTargetNetworkCmd.MarkFlagRequired("association-id")
		ec2_disassociateClientVpnTargetNetworkCmd.MarkFlagRequired("client-vpn-endpoint-id")
		ec2_disassociateClientVpnTargetNetworkCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_disassociateClientVpnTargetNetworkCmd)
}
