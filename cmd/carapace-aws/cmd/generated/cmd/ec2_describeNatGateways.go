package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeNatGatewaysCmd = &cobra.Command{
	Use:   "describe-nat-gateways",
	Short: "Describes your NAT gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeNatGatewaysCmd).Standalone()

	ec2_describeNatGatewaysCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeNatGatewaysCmd.Flags().String("filter", "", "The filters.")
	ec2_describeNatGatewaysCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeNatGatewaysCmd.Flags().String("nat-gateway-ids", "", "The IDs of the NAT gateways.")
	ec2_describeNatGatewaysCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeNatGatewaysCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeNatGatewaysCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeNatGatewaysCmd)
}
