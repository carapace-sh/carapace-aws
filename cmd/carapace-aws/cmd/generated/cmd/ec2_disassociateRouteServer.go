package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateRouteServerCmd = &cobra.Command{
	Use:   "disassociate-route-server",
	Short: "Disassociates a route server from a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateRouteServerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disassociateRouteServerCmd).Standalone()

		ec2_disassociateRouteServerCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_disassociateRouteServerCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_disassociateRouteServerCmd.Flags().String("route-server-id", "", "The ID of the route server to disassociate.")
		ec2_disassociateRouteServerCmd.Flags().String("vpc-id", "", "The ID of the VPC to disassociate from the route server.")
		ec2_disassociateRouteServerCmd.Flag("no-dry-run").Hidden = true
		ec2_disassociateRouteServerCmd.MarkFlagRequired("route-server-id")
		ec2_disassociateRouteServerCmd.MarkFlagRequired("vpc-id")
	})
	ec2Cmd.AddCommand(ec2_disassociateRouteServerCmd)
}
