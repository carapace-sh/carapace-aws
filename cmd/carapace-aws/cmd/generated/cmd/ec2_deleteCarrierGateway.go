package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteCarrierGatewayCmd = &cobra.Command{
	Use:   "delete-carrier-gateway",
	Short: "Deletes a carrier gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteCarrierGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteCarrierGatewayCmd).Standalone()

		ec2_deleteCarrierGatewayCmd.Flags().String("carrier-gateway-id", "", "The ID of the carrier gateway.")
		ec2_deleteCarrierGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteCarrierGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteCarrierGatewayCmd.MarkFlagRequired("carrier-gateway-id")
		ec2_deleteCarrierGatewayCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteCarrierGatewayCmd)
}
