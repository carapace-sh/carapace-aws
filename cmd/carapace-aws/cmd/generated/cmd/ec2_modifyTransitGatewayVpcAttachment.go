package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyTransitGatewayVpcAttachmentCmd = &cobra.Command{
	Use:   "modify-transit-gateway-vpc-attachment",
	Short: "Modifies the specified VPC attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyTransitGatewayVpcAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyTransitGatewayVpcAttachmentCmd).Standalone()

		ec2_modifyTransitGatewayVpcAttachmentCmd.Flags().String("add-subnet-ids", "", "The IDs of one or more subnets to add.")
		ec2_modifyTransitGatewayVpcAttachmentCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyTransitGatewayVpcAttachmentCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyTransitGatewayVpcAttachmentCmd.Flags().String("options", "", "The new VPC attachment options.")
		ec2_modifyTransitGatewayVpcAttachmentCmd.Flags().String("remove-subnet-ids", "", "The IDs of one or more subnets to remove.")
		ec2_modifyTransitGatewayVpcAttachmentCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the attachment.")
		ec2_modifyTransitGatewayVpcAttachmentCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyTransitGatewayVpcAttachmentCmd.MarkFlagRequired("transit-gateway-attachment-id")
	})
	ec2Cmd.AddCommand(ec2_modifyTransitGatewayVpcAttachmentCmd)
}
