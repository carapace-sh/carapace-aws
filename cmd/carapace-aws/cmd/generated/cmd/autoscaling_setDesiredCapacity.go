package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_setDesiredCapacityCmd = &cobra.Command{
	Use:   "set-desired-capacity",
	Short: "Sets the size of the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_setDesiredCapacityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_setDesiredCapacityCmd).Standalone()

		autoscaling_setDesiredCapacityCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_setDesiredCapacityCmd.Flags().String("desired-capacity", "", "The desired capacity is the initial capacity of the Auto Scaling group after this operation completes and the capacity it attempts to maintain.")
		autoscaling_setDesiredCapacityCmd.Flags().String("honor-cooldown", "", "Indicates whether Amazon EC2 Auto Scaling waits for the cooldown period to complete before initiating a scaling activity to set your Auto Scaling group to its new capacity.")
		autoscaling_setDesiredCapacityCmd.MarkFlagRequired("auto-scaling-group-name")
		autoscaling_setDesiredCapacityCmd.MarkFlagRequired("desired-capacity")
	})
	autoscalingCmd.AddCommand(autoscaling_setDesiredCapacityCmd)
}
