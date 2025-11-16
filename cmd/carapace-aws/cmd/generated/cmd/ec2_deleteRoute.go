package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteRouteCmd = &cobra.Command{
	Use:   "delete-route",
	Short: "Deletes the specified route from the specified route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteRouteCmd).Standalone()

		ec2_deleteRouteCmd.Flags().String("destination-cidr-block", "", "The IPv4 CIDR range for the route.")
		ec2_deleteRouteCmd.Flags().String("destination-ipv6-cidr-block", "", "The IPv6 CIDR range for the route.")
		ec2_deleteRouteCmd.Flags().String("destination-prefix-list-id", "", "The ID of the prefix list for the route.")
		ec2_deleteRouteCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteRouteCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteRouteCmd.Flags().String("route-table-id", "", "The ID of the route table.")
		ec2_deleteRouteCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteRouteCmd.MarkFlagRequired("route-table-id")
	})
	ec2Cmd.AddCommand(ec2_deleteRouteCmd)
}
