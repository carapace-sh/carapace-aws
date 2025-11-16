package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateTransitGatewayPolicyTableCmd = &cobra.Command{
	Use:   "associate-transit-gateway-policy-table",
	Short: "Associates the specified transit gateway attachment with a transit gateway policy table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateTransitGatewayPolicyTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_associateTransitGatewayPolicyTableCmd).Standalone()

		ec2_associateTransitGatewayPolicyTableCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_associateTransitGatewayPolicyTableCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_associateTransitGatewayPolicyTableCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the transit gateway attachment to associate with the policy table.")
		ec2_associateTransitGatewayPolicyTableCmd.Flags().String("transit-gateway-policy-table-id", "", "The ID of the transit gateway policy table to associate with the transit gateway attachment.")
		ec2_associateTransitGatewayPolicyTableCmd.Flag("no-dry-run").Hidden = true
		ec2_associateTransitGatewayPolicyTableCmd.MarkFlagRequired("transit-gateway-attachment-id")
		ec2_associateTransitGatewayPolicyTableCmd.MarkFlagRequired("transit-gateway-policy-table-id")
	})
	ec2Cmd.AddCommand(ec2_associateTransitGatewayPolicyTableCmd)
}
