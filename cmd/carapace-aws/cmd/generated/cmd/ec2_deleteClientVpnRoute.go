package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteClientVpnRouteCmd = &cobra.Command{
	Use:   "delete-client-vpn-route",
	Short: "Deletes a route from a Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteClientVpnRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteClientVpnRouteCmd).Standalone()

		ec2_deleteClientVpnRouteCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint from which the route is to be deleted.")
		ec2_deleteClientVpnRouteCmd.Flags().String("destination-cidr-block", "", "The IPv4 address range, in CIDR notation, of the route to be deleted.")
		ec2_deleteClientVpnRouteCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteClientVpnRouteCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteClientVpnRouteCmd.Flags().String("target-vpc-subnet-id", "", "The ID of the target subnet used by the route.")
		ec2_deleteClientVpnRouteCmd.MarkFlagRequired("client-vpn-endpoint-id")
		ec2_deleteClientVpnRouteCmd.MarkFlagRequired("destination-cidr-block")
		ec2_deleteClientVpnRouteCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteClientVpnRouteCmd)
}
