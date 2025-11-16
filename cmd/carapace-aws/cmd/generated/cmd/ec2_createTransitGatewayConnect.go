package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTransitGatewayConnectCmd = &cobra.Command{
	Use:   "create-transit-gateway-connect",
	Short: "Creates a Connect attachment from a specified transit gateway attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTransitGatewayConnectCmd).Standalone()

	ec2_createTransitGatewayConnectCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTransitGatewayConnectCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTransitGatewayConnectCmd.Flags().String("options", "", "The Connect attachment options.")
	ec2_createTransitGatewayConnectCmd.Flags().String("tag-specifications", "", "The tags to apply to the Connect attachment.")
	ec2_createTransitGatewayConnectCmd.Flags().String("transport-transit-gateway-attachment-id", "", "The ID of the transit gateway attachment.")
	ec2_createTransitGatewayConnectCmd.Flag("no-dry-run").Hidden = true
	ec2_createTransitGatewayConnectCmd.MarkFlagRequired("options")
	ec2_createTransitGatewayConnectCmd.MarkFlagRequired("transport-transit-gateway-attachment-id")
	ec2Cmd.AddCommand(ec2_createTransitGatewayConnectCmd)
}
