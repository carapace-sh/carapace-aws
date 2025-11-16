package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeInternetGatewaysCmd = &cobra.Command{
	Use:   "describe-internet-gateways",
	Short: "Describes your internet gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeInternetGatewaysCmd).Standalone()

	ec2_describeInternetGatewaysCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeInternetGatewaysCmd.Flags().String("filters", "", "The filters.")
	ec2_describeInternetGatewaysCmd.Flags().String("internet-gateway-ids", "", "The IDs of the internet gateways.")
	ec2_describeInternetGatewaysCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeInternetGatewaysCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeInternetGatewaysCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeInternetGatewaysCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeInternetGatewaysCmd)
}
