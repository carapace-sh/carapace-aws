package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_searchTransitGatewayRoutesCmd = &cobra.Command{
	Use:   "search-transit-gateway-routes",
	Short: "Searches for routes in the specified transit gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_searchTransitGatewayRoutesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_searchTransitGatewayRoutesCmd).Standalone()

		ec2_searchTransitGatewayRoutesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_searchTransitGatewayRoutesCmd.Flags().String("filters", "", "One or more filters.")
		ec2_searchTransitGatewayRoutesCmd.Flags().String("max-results", "", "The maximum number of routes to return.")
		ec2_searchTransitGatewayRoutesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_searchTransitGatewayRoutesCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the transit gateway route table.")
		ec2_searchTransitGatewayRoutesCmd.MarkFlagRequired("filters")
		ec2_searchTransitGatewayRoutesCmd.Flag("no-dry-run").Hidden = true
		ec2_searchTransitGatewayRoutesCmd.MarkFlagRequired("transit-gateway-route-table-id")
	})
	ec2Cmd.AddCommand(ec2_searchTransitGatewayRoutesCmd)
}
