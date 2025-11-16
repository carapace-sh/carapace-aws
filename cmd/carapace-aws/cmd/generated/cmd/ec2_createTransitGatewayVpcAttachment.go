package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTransitGatewayVpcAttachmentCmd = &cobra.Command{
	Use:   "create-transit-gateway-vpc-attachment",
	Short: "Attaches the specified VPC to the specified transit gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTransitGatewayVpcAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createTransitGatewayVpcAttachmentCmd).Standalone()

		ec2_createTransitGatewayVpcAttachmentCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTransitGatewayVpcAttachmentCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTransitGatewayVpcAttachmentCmd.Flags().String("options", "", "The VPC attachment options.")
		ec2_createTransitGatewayVpcAttachmentCmd.Flags().String("subnet-ids", "", "The IDs of one or more subnets.")
		ec2_createTransitGatewayVpcAttachmentCmd.Flags().String("tag-specifications", "", "The tags to apply to the VPC attachment.")
		ec2_createTransitGatewayVpcAttachmentCmd.Flags().String("transit-gateway-id", "", "The ID of the transit gateway.")
		ec2_createTransitGatewayVpcAttachmentCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
		ec2_createTransitGatewayVpcAttachmentCmd.Flag("no-dry-run").Hidden = true
		ec2_createTransitGatewayVpcAttachmentCmd.MarkFlagRequired("subnet-ids")
		ec2_createTransitGatewayVpcAttachmentCmd.MarkFlagRequired("transit-gateway-id")
		ec2_createTransitGatewayVpcAttachmentCmd.MarkFlagRequired("vpc-id")
	})
	ec2Cmd.AddCommand(ec2_createTransitGatewayVpcAttachmentCmd)
}
