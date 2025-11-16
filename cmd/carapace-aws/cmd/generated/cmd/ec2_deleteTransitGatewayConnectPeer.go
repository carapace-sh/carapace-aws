package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTransitGatewayConnectPeerCmd = &cobra.Command{
	Use:   "delete-transit-gateway-connect-peer",
	Short: "Deletes the specified Connect peer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTransitGatewayConnectPeerCmd).Standalone()

	ec2_deleteTransitGatewayConnectPeerCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteTransitGatewayConnectPeerCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteTransitGatewayConnectPeerCmd.Flags().String("transit-gateway-connect-peer-id", "", "The ID of the Connect peer.")
	ec2_deleteTransitGatewayConnectPeerCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteTransitGatewayConnectPeerCmd.MarkFlagRequired("transit-gateway-connect-peer-id")
	ec2Cmd.AddCommand(ec2_deleteTransitGatewayConnectPeerCmd)
}
