package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeRouteServerPeersCmd = &cobra.Command{
	Use:   "describe-route-server-peers",
	Short: "Describes one or more route server peers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeRouteServerPeersCmd).Standalone()

	ec2_describeRouteServerPeersCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_describeRouteServerPeersCmd.Flags().String("filters", "", "One or more filters to apply to the describe request.")
	ec2_describeRouteServerPeersCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describeRouteServerPeersCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeRouteServerPeersCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_describeRouteServerPeersCmd.Flags().String("route-server-peer-ids", "", "The IDs of the route server peers to describe.")
	ec2_describeRouteServerPeersCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeRouteServerPeersCmd)
}
