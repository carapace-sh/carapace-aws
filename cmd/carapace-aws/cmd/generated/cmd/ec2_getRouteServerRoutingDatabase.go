package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getRouteServerRoutingDatabaseCmd = &cobra.Command{
	Use:   "get-route-server-routing-database",
	Short: "Gets the routing database for the specified route server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getRouteServerRoutingDatabaseCmd).Standalone()

	ec2_getRouteServerRoutingDatabaseCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_getRouteServerRoutingDatabaseCmd.Flags().String("filters", "", "Filters to apply to the routing database query.")
	ec2_getRouteServerRoutingDatabaseCmd.Flags().String("max-results", "", "The maximum number of routing database entries to return in a single response.")
	ec2_getRouteServerRoutingDatabaseCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_getRouteServerRoutingDatabaseCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_getRouteServerRoutingDatabaseCmd.Flags().String("route-server-id", "", "The ID of the route server for which to get the routing database.")
	ec2_getRouteServerRoutingDatabaseCmd.Flag("no-dry-run").Hidden = true
	ec2_getRouteServerRoutingDatabaseCmd.MarkFlagRequired("route-server-id")
	ec2Cmd.AddCommand(ec2_getRouteServerRoutingDatabaseCmd)
}
