package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_putLifecycleHookCmd = &cobra.Command{
	Use:   "put-lifecycle-hook",
	Short: "Creates or updates a lifecycle hook for the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_putLifecycleHookCmd).Standalone()

	autoscaling_putLifecycleHookCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_putLifecycleHookCmd.Flags().String("default-result", "", "The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs.")
	autoscaling_putLifecycleHookCmd.Flags().String("heartbeat-timeout", "", "The maximum time, in seconds, that can elapse before the lifecycle hook times out.")
	autoscaling_putLifecycleHookCmd.Flags().String("lifecycle-hook-name", "", "The name of the lifecycle hook.")
	autoscaling_putLifecycleHookCmd.Flags().String("lifecycle-transition", "", "The lifecycle transition.")
	autoscaling_putLifecycleHookCmd.Flags().String("notification-metadata", "", "Additional information that you want to include any time Amazon EC2 Auto Scaling sends a message to the notification target.")
	autoscaling_putLifecycleHookCmd.Flags().String("notification-target-arn", "", "The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in a wait state for the lifecycle hook.")
	autoscaling_putLifecycleHookCmd.Flags().String("role-arn", "", "The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.")
	autoscaling_putLifecycleHookCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscaling_putLifecycleHookCmd.MarkFlagRequired("lifecycle-hook-name")
	autoscalingCmd.AddCommand(autoscaling_putLifecycleHookCmd)
}
