package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTransitGatewayConnectPeersCmd = &cobra.Command{
	Use:   "describe-transit-gateway-connect-peers",
	Short: "Describes one or more Connect peers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTransitGatewayConnectPeersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeTransitGatewayConnectPeersCmd).Standalone()

		ec2_describeTransitGatewayConnectPeersCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTransitGatewayConnectPeersCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeTransitGatewayConnectPeersCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeTransitGatewayConnectPeersCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeTransitGatewayConnectPeersCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTransitGatewayConnectPeersCmd.Flags().String("transit-gateway-connect-peer-ids", "", "The IDs of the Connect peers.")
		ec2_describeTransitGatewayConnectPeersCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeTransitGatewayConnectPeersCmd)
}
