package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_rejectTransitGatewayPeeringAttachmentCmd = &cobra.Command{
	Use:   "reject-transit-gateway-peering-attachment",
	Short: "Rejects a transit gateway peering attachment request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_rejectTransitGatewayPeeringAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_rejectTransitGatewayPeeringAttachmentCmd).Standalone()

		ec2_rejectTransitGatewayPeeringAttachmentCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_rejectTransitGatewayPeeringAttachmentCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_rejectTransitGatewayPeeringAttachmentCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the transit gateway peering attachment.")
		ec2_rejectTransitGatewayPeeringAttachmentCmd.Flag("no-dry-run").Hidden = true
		ec2_rejectTransitGatewayPeeringAttachmentCmd.MarkFlagRequired("transit-gateway-attachment-id")
	})
	ec2Cmd.AddCommand(ec2_rejectTransitGatewayPeeringAttachmentCmd)
}
