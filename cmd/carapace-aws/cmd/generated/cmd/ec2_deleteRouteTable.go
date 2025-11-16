package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteRouteTableCmd = &cobra.Command{
	Use:   "delete-route-table",
	Short: "Deletes the specified route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteRouteTableCmd).Standalone()

	ec2_deleteRouteTableCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteRouteTableCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteRouteTableCmd.Flags().String("route-table-id", "", "The ID of the route table.")
	ec2_deleteRouteTableCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteRouteTableCmd.MarkFlagRequired("route-table-id")
	ec2Cmd.AddCommand(ec2_deleteRouteTableCmd)
}
