package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeLocalGatewayRouteTablesCmd = &cobra.Command{
	Use:   "describe-local-gateway-route-tables",
	Short: "Describes one or more local gateway route tables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeLocalGatewayRouteTablesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeLocalGatewayRouteTablesCmd).Standalone()

		ec2_describeLocalGatewayRouteTablesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeLocalGatewayRouteTablesCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeLocalGatewayRouteTablesCmd.Flags().String("local-gateway-route-table-ids", "", "The IDs of the local gateway route tables.")
		ec2_describeLocalGatewayRouteTablesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeLocalGatewayRouteTablesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeLocalGatewayRouteTablesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeLocalGatewayRouteTablesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeLocalGatewayRouteTablesCmd)
}
