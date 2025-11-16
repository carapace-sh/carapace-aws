package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpcEndpointConnectionNotificationsCmd = &cobra.Command{
	Use:   "describe-vpc-endpoint-connection-notifications",
	Short: "Describes the connection notifications for VPC endpoints and VPC endpoint services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpcEndpointConnectionNotificationsCmd).Standalone()

	ec2_describeVpcEndpointConnectionNotificationsCmd.Flags().String("connection-notification-id", "", "The ID of the notification.")
	ec2_describeVpcEndpointConnectionNotificationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeVpcEndpointConnectionNotificationsCmd.Flags().String("filters", "", "The filters.")
	ec2_describeVpcEndpointConnectionNotificationsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	ec2_describeVpcEndpointConnectionNotificationsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	ec2_describeVpcEndpointConnectionNotificationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeVpcEndpointConnectionNotificationsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeVpcEndpointConnectionNotificationsCmd)
}
