package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getTransitGatewayPolicyTableEntriesCmd = &cobra.Command{
	Use:   "get-transit-gateway-policy-table-entries",
	Short: "Returns a list of transit gateway policy table entries.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getTransitGatewayPolicyTableEntriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getTransitGatewayPolicyTableEntriesCmd).Standalone()

		ec2_getTransitGatewayPolicyTableEntriesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getTransitGatewayPolicyTableEntriesCmd.Flags().String("filters", "", "The filters associated with the transit gateway policy table.")
		ec2_getTransitGatewayPolicyTableEntriesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_getTransitGatewayPolicyTableEntriesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_getTransitGatewayPolicyTableEntriesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getTransitGatewayPolicyTableEntriesCmd.Flags().String("transit-gateway-policy-table-id", "", "The ID of the transit gateway policy table.")
		ec2_getTransitGatewayPolicyTableEntriesCmd.Flag("no-dry-run").Hidden = true
		ec2_getTransitGatewayPolicyTableEntriesCmd.MarkFlagRequired("transit-gateway-policy-table-id")
	})
	ec2Cmd.AddCommand(ec2_getTransitGatewayPolicyTableEntriesCmd)
}
