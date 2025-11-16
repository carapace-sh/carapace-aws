package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateTransitGatewayPolicyTableCmd = &cobra.Command{
	Use:   "disassociate-transit-gateway-policy-table",
	Short: "Removes the association between an an attachment and a policy table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateTransitGatewayPolicyTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disassociateTransitGatewayPolicyTableCmd).Standalone()

		ec2_disassociateTransitGatewayPolicyTableCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateTransitGatewayPolicyTableCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateTransitGatewayPolicyTableCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the transit gateway attachment to disassociate from the policy table.")
		ec2_disassociateTransitGatewayPolicyTableCmd.Flags().String("transit-gateway-policy-table-id", "", "The ID of the disassociated policy table.")
		ec2_disassociateTransitGatewayPolicyTableCmd.Flag("no-dry-run").Hidden = true
		ec2_disassociateTransitGatewayPolicyTableCmd.MarkFlagRequired("transit-gateway-attachment-id")
		ec2_disassociateTransitGatewayPolicyTableCmd.MarkFlagRequired("transit-gateway-policy-table-id")
	})
	ec2Cmd.AddCommand(ec2_disassociateTransitGatewayPolicyTableCmd)
}
