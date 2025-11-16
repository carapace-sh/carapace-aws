package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVpcEndpointConnectionNotificationCmd = &cobra.Command{
	Use:   "modify-vpc-endpoint-connection-notification",
	Short: "Modifies a connection notification for VPC endpoint or VPC endpoint service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVpcEndpointConnectionNotificationCmd).Standalone()

	ec2_modifyVpcEndpointConnectionNotificationCmd.Flags().String("connection-events", "", "The events for the endpoint.")
	ec2_modifyVpcEndpointConnectionNotificationCmd.Flags().String("connection-notification-arn", "", "The ARN for the SNS topic for the notification.")
	ec2_modifyVpcEndpointConnectionNotificationCmd.Flags().String("connection-notification-id", "", "The ID of the notification.")
	ec2_modifyVpcEndpointConnectionNotificationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVpcEndpointConnectionNotificationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVpcEndpointConnectionNotificationCmd.MarkFlagRequired("connection-notification-id")
	ec2_modifyVpcEndpointConnectionNotificationCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyVpcEndpointConnectionNotificationCmd)
}
