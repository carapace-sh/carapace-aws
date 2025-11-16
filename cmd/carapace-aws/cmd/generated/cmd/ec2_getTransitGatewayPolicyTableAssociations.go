package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getTransitGatewayPolicyTableAssociationsCmd = &cobra.Command{
	Use:   "get-transit-gateway-policy-table-associations",
	Short: "Gets a list of the transit gateway policy table associations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getTransitGatewayPolicyTableAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getTransitGatewayPolicyTableAssociationsCmd).Standalone()

		ec2_getTransitGatewayPolicyTableAssociationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getTransitGatewayPolicyTableAssociationsCmd.Flags().String("filters", "", "The filters associated with the transit gateway policy table.")
		ec2_getTransitGatewayPolicyTableAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_getTransitGatewayPolicyTableAssociationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_getTransitGatewayPolicyTableAssociationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getTransitGatewayPolicyTableAssociationsCmd.Flags().String("transit-gateway-policy-table-id", "", "The ID of the transit gateway policy table.")
		ec2_getTransitGatewayPolicyTableAssociationsCmd.Flag("no-dry-run").Hidden = true
		ec2_getTransitGatewayPolicyTableAssociationsCmd.MarkFlagRequired("transit-gateway-policy-table-id")
	})
	ec2Cmd.AddCommand(ec2_getTransitGatewayPolicyTableAssociationsCmd)
}
