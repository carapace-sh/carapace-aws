package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_deleteNotificationConfigurationCmd = &cobra.Command{
	Use:   "delete-notification-configuration",
	Short: "Deletes the specified notification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_deleteNotificationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_deleteNotificationConfigurationCmd).Standalone()

		autoscaling_deleteNotificationConfigurationCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_deleteNotificationConfigurationCmd.Flags().String("topic-arn", "", "The Amazon Resource Name (ARN) of the Amazon SNS topic.")
		autoscaling_deleteNotificationConfigurationCmd.MarkFlagRequired("auto-scaling-group-name")
		autoscaling_deleteNotificationConfigurationCmd.MarkFlagRequired("topic-arn")
	})
	autoscalingCmd.AddCommand(autoscaling_deleteNotificationConfigurationCmd)
}
