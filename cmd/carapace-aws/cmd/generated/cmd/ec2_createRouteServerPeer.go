package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createRouteServerPeerCmd = &cobra.Command{
	Use:   "create-route-server-peer",
	Short: "Creates a new BGP peer for a specified route server endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createRouteServerPeerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createRouteServerPeerCmd).Standalone()

		ec2_createRouteServerPeerCmd.Flags().String("bgp-options", "", "The BGP options for the peer, including ASN (Autonomous System Number) and BFD (Bidrectional Forwarding Detection) settings.")
		ec2_createRouteServerPeerCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_createRouteServerPeerCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_createRouteServerPeerCmd.Flags().String("peer-address", "", "The IPv4 address of the peer device.")
		ec2_createRouteServerPeerCmd.Flags().String("route-server-endpoint-id", "", "The ID of the route server endpoint for which to create a peer.")
		ec2_createRouteServerPeerCmd.Flags().String("tag-specifications", "", "The tags to apply to the route server peer during creation.")
		ec2_createRouteServerPeerCmd.MarkFlagRequired("bgp-options")
		ec2_createRouteServerPeerCmd.Flag("no-dry-run").Hidden = true
		ec2_createRouteServerPeerCmd.MarkFlagRequired("peer-address")
		ec2_createRouteServerPeerCmd.MarkFlagRequired("route-server-endpoint-id")
	})
	ec2Cmd.AddCommand(ec2_createRouteServerPeerCmd)
}
