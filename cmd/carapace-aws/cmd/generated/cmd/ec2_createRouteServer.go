package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createRouteServerCmd = &cobra.Command{
	Use:   "create-route-server",
	Short: "Creates a new route server to manage dynamic routing in a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createRouteServerCmd).Standalone()

	ec2_createRouteServerCmd.Flags().String("amazon-side-asn", "", "The private Autonomous System Number (ASN) for the Amazon side of the BGP session.")
	ec2_createRouteServerCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
	ec2_createRouteServerCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_createRouteServerCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_createRouteServerCmd.Flags().Bool("no-sns-notifications-enabled", false, "Indicates whether SNS notifications should be enabled for route server events.")
	ec2_createRouteServerCmd.Flags().String("persist-routes", "", "Indicates whether routes should be persisted after all BGP sessions are terminated.")
	ec2_createRouteServerCmd.Flags().String("persist-routes-duration", "", "The number of minutes a route server will wait after BGP is re-established to unpersist the routes in the FIB and RIB.")
	ec2_createRouteServerCmd.Flags().Bool("sns-notifications-enabled", false, "Indicates whether SNS notifications should be enabled for route server events.")
	ec2_createRouteServerCmd.Flags().String("tag-specifications", "", "The tags to apply to the route server during creation.")
	ec2_createRouteServerCmd.MarkFlagRequired("amazon-side-asn")
	ec2_createRouteServerCmd.Flag("no-dry-run").Hidden = true
	ec2_createRouteServerCmd.Flag("no-sns-notifications-enabled").Hidden = true
	ec2Cmd.AddCommand(ec2_createRouteServerCmd)
}
