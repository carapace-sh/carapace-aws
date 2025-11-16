package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTransitGatewayPeeringAttachmentCmd = &cobra.Command{
	Use:   "create-transit-gateway-peering-attachment",
	Short: "Requests a transit gateway peering attachment between the specified transit gateway (requester) and a peer transit gateway (accepter).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTransitGatewayPeeringAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createTransitGatewayPeeringAttachmentCmd).Standalone()

		ec2_createTransitGatewayPeeringAttachmentCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTransitGatewayPeeringAttachmentCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTransitGatewayPeeringAttachmentCmd.Flags().String("options", "", "Requests a transit gateway peering attachment.")
		ec2_createTransitGatewayPeeringAttachmentCmd.Flags().String("peer-account-id", "", "The ID of the Amazon Web Services account that owns the peer transit gateway.")
		ec2_createTransitGatewayPeeringAttachmentCmd.Flags().String("peer-region", "", "The Region where the peer transit gateway is located.")
		ec2_createTransitGatewayPeeringAttachmentCmd.Flags().String("peer-transit-gateway-id", "", "The ID of the peer transit gateway with which to create the peering attachment.")
		ec2_createTransitGatewayPeeringAttachmentCmd.Flags().String("tag-specifications", "", "The tags to apply to the transit gateway peering attachment.")
		ec2_createTransitGatewayPeeringAttachmentCmd.Flags().String("transit-gateway-id", "", "The ID of the transit gateway.")
		ec2_createTransitGatewayPeeringAttachmentCmd.Flag("no-dry-run").Hidden = true
		ec2_createTransitGatewayPeeringAttachmentCmd.MarkFlagRequired("peer-account-id")
		ec2_createTransitGatewayPeeringAttachmentCmd.MarkFlagRequired("peer-region")
		ec2_createTransitGatewayPeeringAttachmentCmd.MarkFlagRequired("peer-transit-gateway-id")
		ec2_createTransitGatewayPeeringAttachmentCmd.MarkFlagRequired("transit-gateway-id")
	})
	ec2Cmd.AddCommand(ec2_createTransitGatewayPeeringAttachmentCmd)
}
