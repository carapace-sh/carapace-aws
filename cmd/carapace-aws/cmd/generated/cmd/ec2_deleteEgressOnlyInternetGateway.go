package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteEgressOnlyInternetGatewayCmd = &cobra.Command{
	Use:   "delete-egress-only-internet-gateway",
	Short: "Deletes an egress-only internet gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteEgressOnlyInternetGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteEgressOnlyInternetGatewayCmd).Standalone()

		ec2_deleteEgressOnlyInternetGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteEgressOnlyInternetGatewayCmd.Flags().String("egress-only-internet-gateway-id", "", "The ID of the egress-only internet gateway.")
		ec2_deleteEgressOnlyInternetGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteEgressOnlyInternetGatewayCmd.MarkFlagRequired("egress-only-internet-gateway-id")
		ec2_deleteEgressOnlyInternetGatewayCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteEgressOnlyInternetGatewayCmd)
}
