package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteVpcEndpointConnectionNotificationsCmd = &cobra.Command{
	Use:   "delete-vpc-endpoint-connection-notifications",
	Short: "Deletes the specified VPC endpoint connection notifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteVpcEndpointConnectionNotificationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteVpcEndpointConnectionNotificationsCmd).Standalone()

		ec2_deleteVpcEndpointConnectionNotificationsCmd.Flags().String("connection-notification-ids", "", "The IDs of the notifications.")
		ec2_deleteVpcEndpointConnectionNotificationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteVpcEndpointConnectionNotificationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteVpcEndpointConnectionNotificationsCmd.MarkFlagRequired("connection-notification-ids")
		ec2_deleteVpcEndpointConnectionNotificationsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteVpcEndpointConnectionNotificationsCmd)
}
