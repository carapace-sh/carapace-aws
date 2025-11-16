package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTransitGatewayPrefixListReferenceCmd = &cobra.Command{
	Use:   "delete-transit-gateway-prefix-list-reference",
	Short: "Deletes a reference (route) to a prefix list in a specified transit gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTransitGatewayPrefixListReferenceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteTransitGatewayPrefixListReferenceCmd).Standalone()

		ec2_deleteTransitGatewayPrefixListReferenceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTransitGatewayPrefixListReferenceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTransitGatewayPrefixListReferenceCmd.Flags().String("prefix-list-id", "", "The ID of the prefix list.")
		ec2_deleteTransitGatewayPrefixListReferenceCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the route table.")
		ec2_deleteTransitGatewayPrefixListReferenceCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteTransitGatewayPrefixListReferenceCmd.MarkFlagRequired("prefix-list-id")
		ec2_deleteTransitGatewayPrefixListReferenceCmd.MarkFlagRequired("transit-gateway-route-table-id")
	})
	ec2Cmd.AddCommand(ec2_deleteTransitGatewayPrefixListReferenceCmd)
}
