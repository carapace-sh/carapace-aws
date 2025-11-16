package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableRouteServerPropagationCmd = &cobra.Command{
	Use:   "enable-route-server-propagation",
	Short: "Defines which route tables the route server can update with routes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableRouteServerPropagationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableRouteServerPropagationCmd).Standalone()

		ec2_enableRouteServerPropagationCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_enableRouteServerPropagationCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_enableRouteServerPropagationCmd.Flags().String("route-server-id", "", "The ID of the route server for which to enable propagation.")
		ec2_enableRouteServerPropagationCmd.Flags().String("route-table-id", "", "The ID of the route table to which route server will propagate routes.")
		ec2_enableRouteServerPropagationCmd.Flag("no-dry-run").Hidden = true
		ec2_enableRouteServerPropagationCmd.MarkFlagRequired("route-server-id")
		ec2_enableRouteServerPropagationCmd.MarkFlagRequired("route-table-id")
	})
	ec2Cmd.AddCommand(ec2_enableRouteServerPropagationCmd)
}
