package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateRouteTableCmd = &cobra.Command{
	Use:   "associate-route-table",
	Short: "Associates a subnet in your VPC or an internet gateway or virtual private gateway attached to your VPC with a route table in your VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateRouteTableCmd).Standalone()

	ec2_associateRouteTableCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateRouteTableCmd.Flags().String("gateway-id", "", "The ID of the internet gateway or virtual private gateway.")
	ec2_associateRouteTableCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateRouteTableCmd.Flags().String("public-ipv4-pool", "", "The ID of a public IPv4 pool.")
	ec2_associateRouteTableCmd.Flags().String("route-table-id", "", "The ID of the route table.")
	ec2_associateRouteTableCmd.Flags().String("subnet-id", "", "The ID of the subnet.")
	ec2_associateRouteTableCmd.Flag("no-dry-run").Hidden = true
	ec2_associateRouteTableCmd.MarkFlagRequired("route-table-id")
	ec2Cmd.AddCommand(ec2_associateRouteTableCmd)
}
