package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_exitStandbyCmd = &cobra.Command{
	Use:   "exit-standby",
	Short: "Moves the specified instances out of the standby state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_exitStandbyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_exitStandbyCmd).Standalone()

		autoscaling_exitStandbyCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_exitStandbyCmd.Flags().String("instance-ids", "", "The IDs of the instances.")
		autoscaling_exitStandbyCmd.MarkFlagRequired("auto-scaling-group-name")
	})
	autoscalingCmd.AddCommand(autoscaling_exitStandbyCmd)
}
