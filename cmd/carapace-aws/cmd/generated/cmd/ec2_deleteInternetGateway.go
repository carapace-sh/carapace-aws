package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteInternetGatewayCmd = &cobra.Command{
	Use:   "delete-internet-gateway",
	Short: "Deletes the specified internet gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteInternetGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteInternetGatewayCmd).Standalone()

		ec2_deleteInternetGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteInternetGatewayCmd.Flags().String("internet-gateway-id", "", "The ID of the internet gateway.")
		ec2_deleteInternetGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteInternetGatewayCmd.MarkFlagRequired("internet-gateway-id")
		ec2_deleteInternetGatewayCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteInternetGatewayCmd)
}
