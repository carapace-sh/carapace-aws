package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeAutoScalingNotificationTypesCmd = &cobra.Command{
	Use:   "describe-auto-scaling-notification-types",
	Short: "Describes the notification types that are supported by Amazon EC2 Auto Scaling.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeAutoScalingNotificationTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_describeAutoScalingNotificationTypesCmd).Standalone()

	})
	autoscalingCmd.AddCommand(autoscaling_describeAutoScalingNotificationTypesCmd)
}
