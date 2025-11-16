package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeCarrierGatewaysCmd = &cobra.Command{
	Use:   "describe-carrier-gateways",
	Short: "Describes one or more of your carrier gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeCarrierGatewaysCmd).Standalone()

	ec2_describeCarrierGatewaysCmd.Flags().String("carrier-gateway-ids", "", "One or more carrier gateway IDs.")
	ec2_describeCarrierGatewaysCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeCarrierGatewaysCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeCarrierGatewaysCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describeCarrierGatewaysCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeCarrierGatewaysCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeCarrierGatewaysCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeCarrierGatewaysCmd)
}
