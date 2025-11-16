package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteRouteServerCmd = &cobra.Command{
	Use:   "delete-route-server",
	Short: "Deletes the specified route server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteRouteServerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteRouteServerCmd).Standalone()

		ec2_deleteRouteServerCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_deleteRouteServerCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_deleteRouteServerCmd.Flags().String("route-server-id", "", "The ID of the route server to delete.")
		ec2_deleteRouteServerCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteRouteServerCmd.MarkFlagRequired("route-server-id")
	})
	ec2Cmd.AddCommand(ec2_deleteRouteServerCmd)
}
