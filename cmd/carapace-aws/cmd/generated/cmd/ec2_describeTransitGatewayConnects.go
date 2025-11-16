package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTransitGatewayConnectsCmd = &cobra.Command{
	Use:   "describe-transit-gateway-connects",
	Short: "Describes one or more Connect attachments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTransitGatewayConnectsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeTransitGatewayConnectsCmd).Standalone()

		ec2_describeTransitGatewayConnectsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTransitGatewayConnectsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeTransitGatewayConnectsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeTransitGatewayConnectsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeTransitGatewayConnectsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTransitGatewayConnectsCmd.Flags().String("transit-gateway-attachment-ids", "", "The IDs of the attachments.")
		ec2_describeTransitGatewayConnectsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeTransitGatewayConnectsCmd)
}
