package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_recordLifecycleActionHeartbeatCmd = &cobra.Command{
	Use:   "record-lifecycle-action-heartbeat",
	Short: "Records a heartbeat for the lifecycle action associated with the specified token or instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_recordLifecycleActionHeartbeatCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_recordLifecycleActionHeartbeatCmd).Standalone()

		autoscaling_recordLifecycleActionHeartbeatCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_recordLifecycleActionHeartbeatCmd.Flags().String("instance-id", "", "The ID of the instance.")
		autoscaling_recordLifecycleActionHeartbeatCmd.Flags().String("lifecycle-action-token", "", "A token that uniquely identifies a specific lifecycle action associated with an instance.")
		autoscaling_recordLifecycleActionHeartbeatCmd.Flags().String("lifecycle-hook-name", "", "The name of the lifecycle hook.")
		autoscaling_recordLifecycleActionHeartbeatCmd.MarkFlagRequired("auto-scaling-group-name")
		autoscaling_recordLifecycleActionHeartbeatCmd.MarkFlagRequired("lifecycle-hook-name")
	})
	autoscalingCmd.AddCommand(autoscaling_recordLifecycleActionHeartbeatCmd)
}
