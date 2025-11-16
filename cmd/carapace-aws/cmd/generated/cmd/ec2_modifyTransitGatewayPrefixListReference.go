package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyTransitGatewayPrefixListReferenceCmd = &cobra.Command{
	Use:   "modify-transit-gateway-prefix-list-reference",
	Short: "Modifies a reference (route) to a prefix list in a specified transit gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyTransitGatewayPrefixListReferenceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyTransitGatewayPrefixListReferenceCmd).Standalone()

		ec2_modifyTransitGatewayPrefixListReferenceCmd.Flags().Bool("blackhole", false, "Indicates whether to drop traffic that matches this route.")
		ec2_modifyTransitGatewayPrefixListReferenceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyTransitGatewayPrefixListReferenceCmd.Flags().Bool("no-blackhole", false, "Indicates whether to drop traffic that matches this route.")
		ec2_modifyTransitGatewayPrefixListReferenceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyTransitGatewayPrefixListReferenceCmd.Flags().String("prefix-list-id", "", "The ID of the prefix list.")
		ec2_modifyTransitGatewayPrefixListReferenceCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the attachment to which traffic is routed.")
		ec2_modifyTransitGatewayPrefixListReferenceCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the transit gateway route table.")
		ec2_modifyTransitGatewayPrefixListReferenceCmd.Flag("no-blackhole").Hidden = true
		ec2_modifyTransitGatewayPrefixListReferenceCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyTransitGatewayPrefixListReferenceCmd.MarkFlagRequired("prefix-list-id")
		ec2_modifyTransitGatewayPrefixListReferenceCmd.MarkFlagRequired("transit-gateway-route-table-id")
	})
	ec2Cmd.AddCommand(ec2_modifyTransitGatewayPrefixListReferenceCmd)
}
