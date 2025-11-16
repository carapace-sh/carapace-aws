package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTransitGatewayPrefixListReferenceCmd = &cobra.Command{
	Use:   "create-transit-gateway-prefix-list-reference",
	Short: "Creates a reference (route) to a prefix list in a specified transit gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTransitGatewayPrefixListReferenceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createTransitGatewayPrefixListReferenceCmd).Standalone()

		ec2_createTransitGatewayPrefixListReferenceCmd.Flags().Bool("blackhole", false, "Indicates whether to drop traffic that matches this route.")
		ec2_createTransitGatewayPrefixListReferenceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTransitGatewayPrefixListReferenceCmd.Flags().Bool("no-blackhole", false, "Indicates whether to drop traffic that matches this route.")
		ec2_createTransitGatewayPrefixListReferenceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTransitGatewayPrefixListReferenceCmd.Flags().String("prefix-list-id", "", "The ID of the prefix list that is used for destination matches.")
		ec2_createTransitGatewayPrefixListReferenceCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the attachment to which traffic is routed.")
		ec2_createTransitGatewayPrefixListReferenceCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the transit gateway route table.")
		ec2_createTransitGatewayPrefixListReferenceCmd.Flag("no-blackhole").Hidden = true
		ec2_createTransitGatewayPrefixListReferenceCmd.Flag("no-dry-run").Hidden = true
		ec2_createTransitGatewayPrefixListReferenceCmd.MarkFlagRequired("prefix-list-id")
		ec2_createTransitGatewayPrefixListReferenceCmd.MarkFlagRequired("transit-gateway-route-table-id")
	})
	ec2Cmd.AddCommand(ec2_createTransitGatewayPrefixListReferenceCmd)
}
