package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyTransitGatewayCmd = &cobra.Command{
	Use:   "modify-transit-gateway",
	Short: "Modifies the specified transit gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyTransitGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyTransitGatewayCmd).Standalone()

		ec2_modifyTransitGatewayCmd.Flags().String("description", "", "The description for the transit gateway.")
		ec2_modifyTransitGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyTransitGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyTransitGatewayCmd.Flags().String("options", "", "The options to modify.")
		ec2_modifyTransitGatewayCmd.Flags().String("transit-gateway-id", "", "The ID of the transit gateway.")
		ec2_modifyTransitGatewayCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyTransitGatewayCmd.MarkFlagRequired("transit-gateway-id")
	})
	ec2Cmd.AddCommand(ec2_modifyTransitGatewayCmd)
}
