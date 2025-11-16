package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeEgressOnlyInternetGatewaysCmd = &cobra.Command{
	Use:   "describe-egress-only-internet-gateways",
	Short: "Describes your egress-only internet gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeEgressOnlyInternetGatewaysCmd).Standalone()

	ec2_describeEgressOnlyInternetGatewaysCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeEgressOnlyInternetGatewaysCmd.Flags().String("egress-only-internet-gateway-ids", "", "The IDs of the egress-only internet gateways.")
	ec2_describeEgressOnlyInternetGatewaysCmd.Flags().String("filters", "", "The filters.")
	ec2_describeEgressOnlyInternetGatewaysCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeEgressOnlyInternetGatewaysCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeEgressOnlyInternetGatewaysCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeEgressOnlyInternetGatewaysCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeEgressOnlyInternetGatewaysCmd)
}
