package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTransitGatewayCmd = &cobra.Command{
	Use:   "delete-transit-gateway",
	Short: "Deletes the specified transit gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTransitGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteTransitGatewayCmd).Standalone()

		ec2_deleteTransitGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTransitGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTransitGatewayCmd.Flags().String("transit-gateway-id", "", "The ID of the transit gateway.")
		ec2_deleteTransitGatewayCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteTransitGatewayCmd.MarkFlagRequired("transit-gateway-id")
	})
	ec2Cmd.AddCommand(ec2_deleteTransitGatewayCmd)
}
