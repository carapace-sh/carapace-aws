package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getTransitGatewayAttachmentPropagationsCmd = &cobra.Command{
	Use:   "get-transit-gateway-attachment-propagations",
	Short: "Lists the route tables to which the specified resource attachment propagates routes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getTransitGatewayAttachmentPropagationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getTransitGatewayAttachmentPropagationsCmd).Standalone()

		ec2_getTransitGatewayAttachmentPropagationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getTransitGatewayAttachmentPropagationsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_getTransitGatewayAttachmentPropagationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_getTransitGatewayAttachmentPropagationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_getTransitGatewayAttachmentPropagationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getTransitGatewayAttachmentPropagationsCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the attachment.")
		ec2_getTransitGatewayAttachmentPropagationsCmd.Flag("no-dry-run").Hidden = true
		ec2_getTransitGatewayAttachmentPropagationsCmd.MarkFlagRequired("transit-gateway-attachment-id")
	})
	ec2Cmd.AddCommand(ec2_getTransitGatewayAttachmentPropagationsCmd)
}
