package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteRouteServerEndpointCmd = &cobra.Command{
	Use:   "delete-route-server-endpoint",
	Short: "Deletes the specified route server endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteRouteServerEndpointCmd).Standalone()

	ec2_deleteRouteServerEndpointCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deleteRouteServerEndpointCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deleteRouteServerEndpointCmd.Flags().String("route-server-endpoint-id", "", "The ID of the route server endpoint to delete.")
	ec2_deleteRouteServerEndpointCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteRouteServerEndpointCmd.MarkFlagRequired("route-server-endpoint-id")
	ec2Cmd.AddCommand(ec2_deleteRouteServerEndpointCmd)
}
