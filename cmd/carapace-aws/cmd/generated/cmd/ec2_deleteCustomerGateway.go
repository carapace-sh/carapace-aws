package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteCustomerGatewayCmd = &cobra.Command{
	Use:   "delete-customer-gateway",
	Short: "Deletes the specified customer gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteCustomerGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteCustomerGatewayCmd).Standalone()

		ec2_deleteCustomerGatewayCmd.Flags().String("customer-gateway-id", "", "The ID of the customer gateway.")
		ec2_deleteCustomerGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteCustomerGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteCustomerGatewayCmd.MarkFlagRequired("customer-gateway-id")
		ec2_deleteCustomerGatewayCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteCustomerGatewayCmd)
}
