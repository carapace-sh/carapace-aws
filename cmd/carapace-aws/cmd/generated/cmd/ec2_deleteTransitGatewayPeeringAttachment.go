package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTransitGatewayPeeringAttachmentCmd = &cobra.Command{
	Use:   "delete-transit-gateway-peering-attachment",
	Short: "Deletes a transit gateway peering attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTransitGatewayPeeringAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteTransitGatewayPeeringAttachmentCmd).Standalone()

		ec2_deleteTransitGatewayPeeringAttachmentCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTransitGatewayPeeringAttachmentCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTransitGatewayPeeringAttachmentCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the transit gateway peering attachment.")
		ec2_deleteTransitGatewayPeeringAttachmentCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteTransitGatewayPeeringAttachmentCmd.MarkFlagRequired("transit-gateway-attachment-id")
	})
	ec2Cmd.AddCommand(ec2_deleteTransitGatewayPeeringAttachmentCmd)
}
