package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeRouteServerEndpointsCmd = &cobra.Command{
	Use:   "describe-route-server-endpoints",
	Short: "Describes one or more route server endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeRouteServerEndpointsCmd).Standalone()

	ec2_describeRouteServerEndpointsCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_describeRouteServerEndpointsCmd.Flags().String("filters", "", "One or more filters to apply to the describe request.")
	ec2_describeRouteServerEndpointsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describeRouteServerEndpointsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeRouteServerEndpointsCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_describeRouteServerEndpointsCmd.Flags().String("route-server-endpoint-ids", "", "The IDs of the route server endpoints to describe.")
	ec2_describeRouteServerEndpointsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeRouteServerEndpointsCmd)
}
