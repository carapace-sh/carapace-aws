package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteVpnGatewayCmd = &cobra.Command{
	Use:   "delete-vpn-gateway",
	Short: "Deletes the specified virtual private gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteVpnGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteVpnGatewayCmd).Standalone()

		ec2_deleteVpnGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteVpnGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteVpnGatewayCmd.Flags().String("vpn-gateway-id", "", "The ID of the virtual private gateway.")
		ec2_deleteVpnGatewayCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteVpnGatewayCmd.MarkFlagRequired("vpn-gateway-id")
	})
	ec2Cmd.AddCommand(ec2_deleteVpnGatewayCmd)
}
