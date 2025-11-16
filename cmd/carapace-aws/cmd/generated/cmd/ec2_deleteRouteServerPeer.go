package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteRouteServerPeerCmd = &cobra.Command{
	Use:   "delete-route-server-peer",
	Short: "Deletes the specified BGP peer from a route server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteRouteServerPeerCmd).Standalone()

	ec2_deleteRouteServerPeerCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deleteRouteServerPeerCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deleteRouteServerPeerCmd.Flags().String("route-server-peer-id", "", "The ID of the route server peer to delete.")
	ec2_deleteRouteServerPeerCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteRouteServerPeerCmd.MarkFlagRequired("route-server-peer-id")
	ec2Cmd.AddCommand(ec2_deleteRouteServerPeerCmd)
}
