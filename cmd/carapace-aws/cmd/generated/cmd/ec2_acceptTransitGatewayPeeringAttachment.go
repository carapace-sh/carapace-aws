package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_acceptTransitGatewayPeeringAttachmentCmd = &cobra.Command{
	Use:   "accept-transit-gateway-peering-attachment",
	Short: "Accepts a transit gateway peering attachment request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_acceptTransitGatewayPeeringAttachmentCmd).Standalone()

	ec2_acceptTransitGatewayPeeringAttachmentCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_acceptTransitGatewayPeeringAttachmentCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_acceptTransitGatewayPeeringAttachmentCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the transit gateway attachment.")
	ec2_acceptTransitGatewayPeeringAttachmentCmd.Flag("no-dry-run").Hidden = true
	ec2_acceptTransitGatewayPeeringAttachmentCmd.MarkFlagRequired("transit-gateway-attachment-id")
	ec2Cmd.AddCommand(ec2_acceptTransitGatewayPeeringAttachmentCmd)
}
