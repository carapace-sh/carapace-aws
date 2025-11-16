package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createRouteServerEndpointCmd = &cobra.Command{
	Use:   "create-route-server-endpoint",
	Short: "Creates a new endpoint for a route server in a specified subnet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createRouteServerEndpointCmd).Standalone()

	ec2_createRouteServerEndpointCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
	ec2_createRouteServerEndpointCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_createRouteServerEndpointCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_createRouteServerEndpointCmd.Flags().String("route-server-id", "", "The ID of the route server for which to create an endpoint.")
	ec2_createRouteServerEndpointCmd.Flags().String("subnet-id", "", "The ID of the subnet in which to create the route server endpoint.")
	ec2_createRouteServerEndpointCmd.Flags().String("tag-specifications", "", "The tags to apply to the route server endpoint during creation.")
	ec2_createRouteServerEndpointCmd.Flag("no-dry-run").Hidden = true
	ec2_createRouteServerEndpointCmd.MarkFlagRequired("route-server-id")
	ec2_createRouteServerEndpointCmd.MarkFlagRequired("subnet-id")
	ec2Cmd.AddCommand(ec2_createRouteServerEndpointCmd)
}
