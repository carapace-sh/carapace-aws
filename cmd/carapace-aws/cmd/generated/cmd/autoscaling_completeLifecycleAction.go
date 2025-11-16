package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_completeLifecycleActionCmd = &cobra.Command{
	Use:   "complete-lifecycle-action",
	Short: "Completes the lifecycle action for the specified token or instance with the specified result.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_completeLifecycleActionCmd).Standalone()

	autoscaling_completeLifecycleActionCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_completeLifecycleActionCmd.Flags().String("instance-id", "", "The ID of the instance.")
	autoscaling_completeLifecycleActionCmd.Flags().String("lifecycle-action-result", "", "The action for the group to take.")
	autoscaling_completeLifecycleActionCmd.Flags().String("lifecycle-action-token", "", "A universally unique identifier (UUID) that identifies a specific lifecycle action associated with an instance.")
	autoscaling_completeLifecycleActionCmd.Flags().String("lifecycle-hook-name", "", "The name of the lifecycle hook.")
	autoscaling_completeLifecycleActionCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscaling_completeLifecycleActionCmd.MarkFlagRequired("lifecycle-action-result")
	autoscaling_completeLifecycleActionCmd.MarkFlagRequired("lifecycle-hook-name")
	autoscalingCmd.AddCommand(autoscaling_completeLifecycleActionCmd)
}
