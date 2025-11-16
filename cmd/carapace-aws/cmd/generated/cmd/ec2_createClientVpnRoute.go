package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createClientVpnRouteCmd = &cobra.Command{
	Use:   "create-client-vpn-route",
	Short: "Adds a route to a network to a Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createClientVpnRouteCmd).Standalone()

	ec2_createClientVpnRouteCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createClientVpnRouteCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint to which to add the route.")
	ec2_createClientVpnRouteCmd.Flags().String("description", "", "A brief description of the route.")
	ec2_createClientVpnRouteCmd.Flags().String("destination-cidr-block", "", "The IPv4 address range, in CIDR notation, of the route destination.")
	ec2_createClientVpnRouteCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createClientVpnRouteCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createClientVpnRouteCmd.Flags().String("target-vpc-subnet-id", "", "The ID of the subnet through which you want to route traffic.")
	ec2_createClientVpnRouteCmd.MarkFlagRequired("client-vpn-endpoint-id")
	ec2_createClientVpnRouteCmd.MarkFlagRequired("destination-cidr-block")
	ec2_createClientVpnRouteCmd.Flag("no-dry-run").Hidden = true
	ec2_createClientVpnRouteCmd.MarkFlagRequired("target-vpc-subnet-id")
	ec2Cmd.AddCommand(ec2_createClientVpnRouteCmd)
}
