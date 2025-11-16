package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTransitGatewayAttachmentsCmd = &cobra.Command{
	Use:   "describe-transit-gateway-attachments",
	Short: "Describes one or more attachments between resources and transit gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTransitGatewayAttachmentsCmd).Standalone()

	ec2_describeTransitGatewayAttachmentsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeTransitGatewayAttachmentsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeTransitGatewayAttachmentsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describeTransitGatewayAttachmentsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeTransitGatewayAttachmentsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeTransitGatewayAttachmentsCmd.Flags().String("transit-gateway-attachment-ids", "", "The IDs of the attachments.")
	ec2_describeTransitGatewayAttachmentsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeTransitGatewayAttachmentsCmd)
}
