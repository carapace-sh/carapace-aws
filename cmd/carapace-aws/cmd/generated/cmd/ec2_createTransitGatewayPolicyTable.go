package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTransitGatewayPolicyTableCmd = &cobra.Command{
	Use:   "create-transit-gateway-policy-table",
	Short: "Creates a transit gateway policy table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTransitGatewayPolicyTableCmd).Standalone()

	ec2_createTransitGatewayPolicyTableCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTransitGatewayPolicyTableCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTransitGatewayPolicyTableCmd.Flags().String("tag-specifications", "", "The tags specification for the transit gateway policy table created during the request.")
	ec2_createTransitGatewayPolicyTableCmd.Flags().String("transit-gateway-id", "", "The ID of the transit gateway used for the policy table.")
	ec2_createTransitGatewayPolicyTableCmd.Flag("no-dry-run").Hidden = true
	ec2_createTransitGatewayPolicyTableCmd.MarkFlagRequired("transit-gateway-id")
	ec2Cmd.AddCommand(ec2_createTransitGatewayPolicyTableCmd)
}
