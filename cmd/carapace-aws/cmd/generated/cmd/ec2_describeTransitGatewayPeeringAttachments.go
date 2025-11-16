package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTransitGatewayPeeringAttachmentsCmd = &cobra.Command{
	Use:   "describe-transit-gateway-peering-attachments",
	Short: "Describes your transit gateway peering attachments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTransitGatewayPeeringAttachmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeTransitGatewayPeeringAttachmentsCmd).Standalone()

		ec2_describeTransitGatewayPeeringAttachmentsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTransitGatewayPeeringAttachmentsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeTransitGatewayPeeringAttachmentsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeTransitGatewayPeeringAttachmentsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeTransitGatewayPeeringAttachmentsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTransitGatewayPeeringAttachmentsCmd.Flags().String("transit-gateway-attachment-ids", "", "One or more IDs of the transit gateway peering attachments.")
		ec2_describeTransitGatewayPeeringAttachmentsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeTransitGatewayPeeringAttachmentsCmd)
}
