package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createVpnGatewayCmd = &cobra.Command{
	Use:   "create-vpn-gateway",
	Short: "Creates a virtual private gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createVpnGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createVpnGatewayCmd).Standalone()

		ec2_createVpnGatewayCmd.Flags().String("amazon-side-asn", "", "A private Autonomous System Number (ASN) for the Amazon side of a BGP session.")
		ec2_createVpnGatewayCmd.Flags().String("availability-zone", "", "The Availability Zone for the virtual private gateway.")
		ec2_createVpnGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createVpnGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createVpnGatewayCmd.Flags().String("tag-specifications", "", "The tags to apply to the virtual private gateway.")
		ec2_createVpnGatewayCmd.Flags().String("type", "", "The type of VPN connection this virtual private gateway supports.")
		ec2_createVpnGatewayCmd.Flag("no-dry-run").Hidden = true
		ec2_createVpnGatewayCmd.MarkFlagRequired("type")
	})
	ec2Cmd.AddCommand(ec2_createVpnGatewayCmd)
}
