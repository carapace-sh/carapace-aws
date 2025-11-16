package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTransitGatewayConnectCmd = &cobra.Command{
	Use:   "delete-transit-gateway-connect",
	Short: "Deletes the specified Connect attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTransitGatewayConnectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteTransitGatewayConnectCmd).Standalone()

		ec2_deleteTransitGatewayConnectCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTransitGatewayConnectCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTransitGatewayConnectCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the Connect attachment.")
		ec2_deleteTransitGatewayConnectCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteTransitGatewayConnectCmd.MarkFlagRequired("transit-gateway-attachment-id")
	})
	ec2Cmd.AddCommand(ec2_deleteTransitGatewayConnectCmd)
}
