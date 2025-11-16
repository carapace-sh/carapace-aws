package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateRouteServerCmd = &cobra.Command{
	Use:   "associate-route-server",
	Short: "Associates a route server with a VPC to enable dynamic route updates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateRouteServerCmd).Standalone()

	ec2_associateRouteServerCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_associateRouteServerCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_associateRouteServerCmd.Flags().String("route-server-id", "", "The unique identifier for the route server to be associated.")
	ec2_associateRouteServerCmd.Flags().String("vpc-id", "", "The ID of the VPC to associate with the route server.")
	ec2_associateRouteServerCmd.Flag("no-dry-run").Hidden = true
	ec2_associateRouteServerCmd.MarkFlagRequired("route-server-id")
	ec2_associateRouteServerCmd.MarkFlagRequired("vpc-id")
	ec2Cmd.AddCommand(ec2_associateRouteServerCmd)
}
