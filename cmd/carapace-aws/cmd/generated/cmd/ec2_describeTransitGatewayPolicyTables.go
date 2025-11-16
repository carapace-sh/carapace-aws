package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTransitGatewayPolicyTablesCmd = &cobra.Command{
	Use:   "describe-transit-gateway-policy-tables",
	Short: "Describes one or more transit gateway route policy tables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTransitGatewayPolicyTablesCmd).Standalone()

	ec2_describeTransitGatewayPolicyTablesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeTransitGatewayPolicyTablesCmd.Flags().String("filters", "", "The filters associated with the transit gateway policy table.")
	ec2_describeTransitGatewayPolicyTablesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describeTransitGatewayPolicyTablesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeTransitGatewayPolicyTablesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeTransitGatewayPolicyTablesCmd.Flags().String("transit-gateway-policy-table-ids", "", "The IDs of the transit gateway policy tables.")
	ec2_describeTransitGatewayPolicyTablesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeTransitGatewayPolicyTablesCmd)
}
