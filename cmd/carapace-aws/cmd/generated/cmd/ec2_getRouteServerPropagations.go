package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getRouteServerPropagationsCmd = &cobra.Command{
	Use:   "get-route-server-propagations",
	Short: "Gets information about the route propagations for the specified route server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getRouteServerPropagationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getRouteServerPropagationsCmd).Standalone()

		ec2_getRouteServerPropagationsCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getRouteServerPropagationsCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getRouteServerPropagationsCmd.Flags().String("route-server-id", "", "The ID of the route server for which to get propagation information.")
		ec2_getRouteServerPropagationsCmd.Flags().String("route-table-id", "", "The ID of the route table for which to get propagation information.")
		ec2_getRouteServerPropagationsCmd.Flag("no-dry-run").Hidden = true
		ec2_getRouteServerPropagationsCmd.MarkFlagRequired("route-server-id")
	})
	ec2Cmd.AddCommand(ec2_getRouteServerPropagationsCmd)
}
