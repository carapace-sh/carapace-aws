package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_rejectTransitGatewayVpcAttachmentCmd = &cobra.Command{
	Use:   "reject-transit-gateway-vpc-attachment",
	Short: "Rejects a request to attach a VPC to a transit gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_rejectTransitGatewayVpcAttachmentCmd).Standalone()

	ec2_rejectTransitGatewayVpcAttachmentCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_rejectTransitGatewayVpcAttachmentCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_rejectTransitGatewayVpcAttachmentCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the attachment.")
	ec2_rejectTransitGatewayVpcAttachmentCmd.Flag("no-dry-run").Hidden = true
	ec2_rejectTransitGatewayVpcAttachmentCmd.MarkFlagRequired("transit-gateway-attachment-id")
	ec2Cmd.AddCommand(ec2_rejectTransitGatewayVpcAttachmentCmd)
}
