package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getRouteServerAssociationsCmd = &cobra.Command{
	Use:   "get-route-server-associations",
	Short: "Gets information about the associations for the specified route server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getRouteServerAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getRouteServerAssociationsCmd).Standalone()

		ec2_getRouteServerAssociationsCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getRouteServerAssociationsCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getRouteServerAssociationsCmd.Flags().String("route-server-id", "", "The ID of the route server for which to get association information.")
		ec2_getRouteServerAssociationsCmd.Flag("no-dry-run").Hidden = true
		ec2_getRouteServerAssociationsCmd.MarkFlagRequired("route-server-id")
	})
	ec2Cmd.AddCommand(ec2_getRouteServerAssociationsCmd)
}
