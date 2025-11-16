package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTransitGatewayPolicyTableCmd = &cobra.Command{
	Use:   "delete-transit-gateway-policy-table",
	Short: "Deletes the specified transit gateway policy table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTransitGatewayPolicyTableCmd).Standalone()

	ec2_deleteTransitGatewayPolicyTableCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteTransitGatewayPolicyTableCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteTransitGatewayPolicyTableCmd.Flags().String("transit-gateway-policy-table-id", "", "The transit gateway policy table to delete.")
	ec2_deleteTransitGatewayPolicyTableCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteTransitGatewayPolicyTableCmd.MarkFlagRequired("transit-gateway-policy-table-id")
	ec2Cmd.AddCommand(ec2_deleteTransitGatewayPolicyTableCmd)
}
