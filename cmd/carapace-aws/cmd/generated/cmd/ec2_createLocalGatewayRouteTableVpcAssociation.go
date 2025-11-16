package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createLocalGatewayRouteTableVpcAssociationCmd = &cobra.Command{
	Use:   "create-local-gateway-route-table-vpc-association",
	Short: "Associates the specified VPC with the specified local gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createLocalGatewayRouteTableVpcAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createLocalGatewayRouteTableVpcAssociationCmd).Standalone()

		ec2_createLocalGatewayRouteTableVpcAssociationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createLocalGatewayRouteTableVpcAssociationCmd.Flags().String("local-gateway-route-table-id", "", "The ID of the local gateway route table.")
		ec2_createLocalGatewayRouteTableVpcAssociationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createLocalGatewayRouteTableVpcAssociationCmd.Flags().String("tag-specifications", "", "The tags to assign to the local gateway route table VPC association.")
		ec2_createLocalGatewayRouteTableVpcAssociationCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
		ec2_createLocalGatewayRouteTableVpcAssociationCmd.MarkFlagRequired("local-gateway-route-table-id")
		ec2_createLocalGatewayRouteTableVpcAssociationCmd.Flag("no-dry-run").Hidden = true
		ec2_createLocalGatewayRouteTableVpcAssociationCmd.MarkFlagRequired("vpc-id")
	})
	ec2Cmd.AddCommand(ec2_createLocalGatewayRouteTableVpcAssociationCmd)
}
