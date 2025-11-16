package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_startInstanceRefreshCmd = &cobra.Command{
	Use:   "start-instance-refresh",
	Short: "Starts an instance refresh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_startInstanceRefreshCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_startInstanceRefreshCmd).Standalone()

		autoscaling_startInstanceRefreshCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_startInstanceRefreshCmd.Flags().String("desired-configuration", "", "The desired configuration.")
		autoscaling_startInstanceRefreshCmd.Flags().String("preferences", "", "Sets your preferences for the instance refresh so that it performs as expected when you start it.")
		autoscaling_startInstanceRefreshCmd.Flags().String("strategy", "", "The strategy to use for the instance refresh.")
		autoscaling_startInstanceRefreshCmd.MarkFlagRequired("auto-scaling-group-name")
	})
	autoscalingCmd.AddCommand(autoscaling_startInstanceRefreshCmd)
}
