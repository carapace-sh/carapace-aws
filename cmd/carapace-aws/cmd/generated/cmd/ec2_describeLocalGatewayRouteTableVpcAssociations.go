package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeLocalGatewayRouteTableVpcAssociationsCmd = &cobra.Command{
	Use:   "describe-local-gateway-route-table-vpc-associations",
	Short: "Describes the specified associations between VPCs and local gateway route tables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeLocalGatewayRouteTableVpcAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeLocalGatewayRouteTableVpcAssociationsCmd).Standalone()

		ec2_describeLocalGatewayRouteTableVpcAssociationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeLocalGatewayRouteTableVpcAssociationsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeLocalGatewayRouteTableVpcAssociationsCmd.Flags().String("local-gateway-route-table-vpc-association-ids", "", "The IDs of the associations.")
		ec2_describeLocalGatewayRouteTableVpcAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeLocalGatewayRouteTableVpcAssociationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeLocalGatewayRouteTableVpcAssociationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeLocalGatewayRouteTableVpcAssociationsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeLocalGatewayRouteTableVpcAssociationsCmd)
}
