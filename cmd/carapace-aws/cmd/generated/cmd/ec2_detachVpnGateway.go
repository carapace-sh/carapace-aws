package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_detachVpnGatewayCmd = &cobra.Command{
	Use:   "detach-vpn-gateway",
	Short: "Detaches a virtual private gateway from a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_detachVpnGatewayCmd).Standalone()

	ec2_detachVpnGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_detachVpnGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_detachVpnGatewayCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
	ec2_detachVpnGatewayCmd.Flags().String("vpn-gateway-id", "", "The ID of the virtual private gateway.")
	ec2_detachVpnGatewayCmd.Flag("no-dry-run").Hidden = true
	ec2_detachVpnGatewayCmd.MarkFlagRequired("vpc-id")
	ec2_detachVpnGatewayCmd.MarkFlagRequired("vpn-gateway-id")
	ec2Cmd.AddCommand(ec2_detachVpnGatewayCmd)
}
