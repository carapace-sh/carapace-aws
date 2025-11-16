package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_deleteAutoScalingGroupCmd = &cobra.Command{
	Use:   "delete-auto-scaling-group",
	Short: "Deletes the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_deleteAutoScalingGroupCmd).Standalone()

	autoscaling_deleteAutoScalingGroupCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_deleteAutoScalingGroupCmd.Flags().String("force-delete", "", "Specifies that the group is to be deleted along with all instances associated with the group, without waiting for all instances to be terminated.")
	autoscaling_deleteAutoScalingGroupCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscalingCmd.AddCommand(autoscaling_deleteAutoScalingGroupCmd)
}
