package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createVpcEndpointConnectionNotificationCmd = &cobra.Command{
	Use:   "create-vpc-endpoint-connection-notification",
	Short: "Creates a connection notification for a specified VPC endpoint or VPC endpoint service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createVpcEndpointConnectionNotificationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createVpcEndpointConnectionNotificationCmd).Standalone()

		ec2_createVpcEndpointConnectionNotificationCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createVpcEndpointConnectionNotificationCmd.Flags().String("connection-events", "", "The endpoint events for which to receive notifications.")
		ec2_createVpcEndpointConnectionNotificationCmd.Flags().String("connection-notification-arn", "", "The ARN of the SNS topic for the notifications.")
		ec2_createVpcEndpointConnectionNotificationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createVpcEndpointConnectionNotificationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createVpcEndpointConnectionNotificationCmd.Flags().String("service-id", "", "The ID of the endpoint service.")
		ec2_createVpcEndpointConnectionNotificationCmd.Flags().String("vpc-endpoint-id", "", "The ID of the endpoint.")
		ec2_createVpcEndpointConnectionNotificationCmd.MarkFlagRequired("connection-events")
		ec2_createVpcEndpointConnectionNotificationCmd.MarkFlagRequired("connection-notification-arn")
		ec2_createVpcEndpointConnectionNotificationCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createVpcEndpointConnectionNotificationCmd)
}
