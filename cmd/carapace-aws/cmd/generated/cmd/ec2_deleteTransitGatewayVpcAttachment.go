package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTransitGatewayVpcAttachmentCmd = &cobra.Command{
	Use:   "delete-transit-gateway-vpc-attachment",
	Short: "Deletes the specified VPC attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTransitGatewayVpcAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteTransitGatewayVpcAttachmentCmd).Standalone()

		ec2_deleteTransitGatewayVpcAttachmentCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTransitGatewayVpcAttachmentCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTransitGatewayVpcAttachmentCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the attachment.")
		ec2_deleteTransitGatewayVpcAttachmentCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteTransitGatewayVpcAttachmentCmd.MarkFlagRequired("transit-gateway-attachment-id")
	})
	ec2Cmd.AddCommand(ec2_deleteTransitGatewayVpcAttachmentCmd)
}
