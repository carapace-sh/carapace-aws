package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteLocalGatewayRouteTableVpcAssociationCmd = &cobra.Command{
	Use:   "delete-local-gateway-route-table-vpc-association",
	Short: "Deletes the specified association between a VPC and local gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteLocalGatewayRouteTableVpcAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteLocalGatewayRouteTableVpcAssociationCmd).Standalone()

		ec2_deleteLocalGatewayRouteTableVpcAssociationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteLocalGatewayRouteTableVpcAssociationCmd.Flags().String("local-gateway-route-table-vpc-association-id", "", "The ID of the association.")
		ec2_deleteLocalGatewayRouteTableVpcAssociationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteLocalGatewayRouteTableVpcAssociationCmd.MarkFlagRequired("local-gateway-route-table-vpc-association-id")
		ec2_deleteLocalGatewayRouteTableVpcAssociationCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteLocalGatewayRouteTableVpcAssociationCmd)
}
