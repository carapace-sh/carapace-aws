package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTransitGatewayRouteTablesCmd = &cobra.Command{
	Use:   "describe-transit-gateway-route-tables",
	Short: "Describes one or more transit gateway route tables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTransitGatewayRouteTablesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeTransitGatewayRouteTablesCmd).Standalone()

		ec2_describeTransitGatewayRouteTablesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTransitGatewayRouteTablesCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeTransitGatewayRouteTablesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeTransitGatewayRouteTablesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeTransitGatewayRouteTablesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTransitGatewayRouteTablesCmd.Flags().String("transit-gateway-route-table-ids", "", "The IDs of the transit gateway route tables.")
		ec2_describeTransitGatewayRouteTablesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeTransitGatewayRouteTablesCmd)
}
