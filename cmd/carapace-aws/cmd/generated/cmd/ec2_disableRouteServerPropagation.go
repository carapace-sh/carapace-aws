package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableRouteServerPropagationCmd = &cobra.Command{
	Use:   "disable-route-server-propagation",
	Short: "Disables route propagation from a route server to a specified route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableRouteServerPropagationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disableRouteServerPropagationCmd).Standalone()

		ec2_disableRouteServerPropagationCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_disableRouteServerPropagationCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_disableRouteServerPropagationCmd.Flags().String("route-server-id", "", "The ID of the route server for which to disable propagation.")
		ec2_disableRouteServerPropagationCmd.Flags().String("route-table-id", "", "The ID of the route table for which to disable route server propagation.")
		ec2_disableRouteServerPropagationCmd.Flag("no-dry-run").Hidden = true
		ec2_disableRouteServerPropagationCmd.MarkFlagRequired("route-server-id")
		ec2_disableRouteServerPropagationCmd.MarkFlagRequired("route-table-id")
	})
	ec2Cmd.AddCommand(ec2_disableRouteServerPropagationCmd)
}
