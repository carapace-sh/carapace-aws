package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_acceptTransitGatewayVpcAttachmentCmd = &cobra.Command{
	Use:   "accept-transit-gateway-vpc-attachment",
	Short: "Accepts a request to attach a VPC to a transit gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_acceptTransitGatewayVpcAttachmentCmd).Standalone()

	ec2_acceptTransitGatewayVpcAttachmentCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_acceptTransitGatewayVpcAttachmentCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_acceptTransitGatewayVpcAttachmentCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the attachment.")
	ec2_acceptTransitGatewayVpcAttachmentCmd.Flag("no-dry-run").Hidden = true
	ec2_acceptTransitGatewayVpcAttachmentCmd.MarkFlagRequired("transit-gateway-attachment-id")
	ec2Cmd.AddCommand(ec2_acceptTransitGatewayVpcAttachmentCmd)
}
