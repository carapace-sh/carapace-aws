package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_terminateInstanceInAutoScalingGroupCmd = &cobra.Command{
	Use:   "terminate-instance-in-auto-scaling-group",
	Short: "Terminates the specified instance and optionally adjusts the desired group size.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_terminateInstanceInAutoScalingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_terminateInstanceInAutoScalingGroupCmd).Standalone()

		autoscaling_terminateInstanceInAutoScalingGroupCmd.Flags().String("instance-id", "", "The ID of the instance.")
		autoscaling_terminateInstanceInAutoScalingGroupCmd.Flags().String("should-decrement-desired-capacity", "", "Indicates whether terminating the instance also decrements the size of the Auto Scaling group.")
		autoscaling_terminateInstanceInAutoScalingGroupCmd.MarkFlagRequired("instance-id")
		autoscaling_terminateInstanceInAutoScalingGroupCmd.MarkFlagRequired("should-decrement-desired-capacity")
	})
	autoscalingCmd.AddCommand(autoscaling_terminateInstanceInAutoScalingGroupCmd)
}
