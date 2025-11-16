package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyRouteServerCmd = &cobra.Command{
	Use:   "modify-route-server",
	Short: "Modifies the configuration of an existing route server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyRouteServerCmd).Standalone()

	ec2_modifyRouteServerCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_modifyRouteServerCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_modifyRouteServerCmd.Flags().Bool("no-sns-notifications-enabled", false, "Specifies whether to enable SNS notifications for route server events.")
	ec2_modifyRouteServerCmd.Flags().String("persist-routes", "", "Specifies whether to persist routes after all BGP sessions are terminated.")
	ec2_modifyRouteServerCmd.Flags().String("persist-routes-duration", "", "The number of minutes a route server will wait after BGP is re-established to unpersist the routes in the FIB and RIB.")
	ec2_modifyRouteServerCmd.Flags().String("route-server-id", "", "The ID of the route server to modify.")
	ec2_modifyRouteServerCmd.Flags().Bool("sns-notifications-enabled", false, "Specifies whether to enable SNS notifications for route server events.")
	ec2_modifyRouteServerCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyRouteServerCmd.Flag("no-sns-notifications-enabled").Hidden = true
	ec2_modifyRouteServerCmd.MarkFlagRequired("route-server-id")
	ec2Cmd.AddCommand(ec2_modifyRouteServerCmd)
}
