package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTransitGatewayVpcAttachmentsCmd = &cobra.Command{
	Use:   "describe-transit-gateway-vpc-attachments",
	Short: "Describes one or more VPC attachments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTransitGatewayVpcAttachmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeTransitGatewayVpcAttachmentsCmd).Standalone()

		ec2_describeTransitGatewayVpcAttachmentsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTransitGatewayVpcAttachmentsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeTransitGatewayVpcAttachmentsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeTransitGatewayVpcAttachmentsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeTransitGatewayVpcAttachmentsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTransitGatewayVpcAttachmentsCmd.Flags().String("transit-gateway-attachment-ids", "", "The IDs of the attachments.")
		ec2_describeTransitGatewayVpcAttachmentsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeTransitGatewayVpcAttachmentsCmd)
}
