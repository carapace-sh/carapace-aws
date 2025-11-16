package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_enterStandbyCmd = &cobra.Command{
	Use:   "enter-standby",
	Short: "Moves the specified instances into the standby state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_enterStandbyCmd).Standalone()

	autoscaling_enterStandbyCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_enterStandbyCmd.Flags().String("instance-ids", "", "The IDs of the instances.")
	autoscaling_enterStandbyCmd.Flags().String("should-decrement-desired-capacity", "", "Indicates whether to decrement the desired capacity of the Auto Scaling group by the number of instances moved to `Standby` mode.")
	autoscaling_enterStandbyCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscaling_enterStandbyCmd.MarkFlagRequired("should-decrement-desired-capacity")
	autoscalingCmd.AddCommand(autoscaling_enterStandbyCmd)
}
