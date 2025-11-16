package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeCustomerGatewaysCmd = &cobra.Command{
	Use:   "describe-customer-gateways",
	Short: "Describes one or more of your VPN customer gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeCustomerGatewaysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeCustomerGatewaysCmd).Standalone()

		ec2_describeCustomerGatewaysCmd.Flags().String("customer-gateway-ids", "", "One or more customer gateway IDs.")
		ec2_describeCustomerGatewaysCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCustomerGatewaysCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeCustomerGatewaysCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCustomerGatewaysCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeCustomerGatewaysCmd)
}
