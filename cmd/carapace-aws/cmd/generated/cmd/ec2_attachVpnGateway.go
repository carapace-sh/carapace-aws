package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_attachVpnGatewayCmd = &cobra.Command{
	Use:   "attach-vpn-gateway",
	Short: "Attaches an available virtual private gateway to a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_attachVpnGatewayCmd).Standalone()

	ec2_attachVpnGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_attachVpnGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_attachVpnGatewayCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
	ec2_attachVpnGatewayCmd.Flags().String("vpn-gateway-id", "", "The ID of the virtual private gateway.")
	ec2_attachVpnGatewayCmd.Flag("no-dry-run").Hidden = true
	ec2_attachVpnGatewayCmd.MarkFlagRequired("vpc-id")
	ec2_attachVpnGatewayCmd.MarkFlagRequired("vpn-gateway-id")
	ec2Cmd.AddCommand(ec2_attachVpnGatewayCmd)
}
