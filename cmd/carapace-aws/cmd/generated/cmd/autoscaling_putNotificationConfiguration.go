package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_putNotificationConfigurationCmd = &cobra.Command{
	Use:   "put-notification-configuration",
	Short: "Configures an Auto Scaling group to send notifications when specified events take place.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_putNotificationConfigurationCmd).Standalone()

	autoscaling_putNotificationConfigurationCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_putNotificationConfigurationCmd.Flags().String("notification-types", "", "The type of event that causes the notification to be sent.")
	autoscaling_putNotificationConfigurationCmd.Flags().String("topic-arn", "", "The Amazon Resource Name (ARN) of the Amazon SNS topic.")
	autoscaling_putNotificationConfigurationCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscaling_putNotificationConfigurationCmd.MarkFlagRequired("notification-types")
	autoscaling_putNotificationConfigurationCmd.MarkFlagRequired("topic-arn")
	autoscalingCmd.AddCommand(autoscaling_putNotificationConfigurationCmd)
}
