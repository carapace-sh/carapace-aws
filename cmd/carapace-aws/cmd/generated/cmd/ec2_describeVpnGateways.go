package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpnGatewaysCmd = &cobra.Command{
	Use:   "describe-vpn-gateways",
	Short: "Describes one or more of your virtual private gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpnGatewaysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVpnGatewaysCmd).Standalone()

		ec2_describeVpnGatewaysCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpnGatewaysCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeVpnGatewaysCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpnGatewaysCmd.Flags().String("vpn-gateway-ids", "", "One or more virtual private gateway IDs.")
		ec2_describeVpnGatewaysCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeVpnGatewaysCmd)
}
