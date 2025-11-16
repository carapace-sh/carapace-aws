package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_detachInstancesCmd = &cobra.Command{
	Use:   "detach-instances",
	Short: "Removes one or more instances from the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_detachInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_detachInstancesCmd).Standalone()

		autoscaling_detachInstancesCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_detachInstancesCmd.Flags().String("instance-ids", "", "The IDs of the instances.")
		autoscaling_detachInstancesCmd.Flags().String("should-decrement-desired-capacity", "", "Indicates whether the Auto Scaling group decrements the desired capacity value by the number of instances detached.")
		autoscaling_detachInstancesCmd.MarkFlagRequired("auto-scaling-group-name")
		autoscaling_detachInstancesCmd.MarkFlagRequired("should-decrement-desired-capacity")
	})
	autoscalingCmd.AddCommand(autoscaling_detachInstancesCmd)
}
