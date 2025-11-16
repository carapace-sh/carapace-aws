package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteNatGatewayCmd = &cobra.Command{
	Use:   "delete-nat-gateway",
	Short: "Deletes the specified NAT gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteNatGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteNatGatewayCmd).Standalone()

		ec2_deleteNatGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteNatGatewayCmd.Flags().String("nat-gateway-id", "", "The ID of the NAT gateway.")
		ec2_deleteNatGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteNatGatewayCmd.MarkFlagRequired("nat-gateway-id")
		ec2_deleteNatGatewayCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteNatGatewayCmd)
}
