package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_disableMetricsCollectionCmd = &cobra.Command{
	Use:   "disable-metrics-collection",
	Short: "Disables group metrics collection for the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_disableMetricsCollectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_disableMetricsCollectionCmd).Standalone()

		autoscaling_disableMetricsCollectionCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_disableMetricsCollectionCmd.Flags().String("metrics", "", "Identifies the metrics to disable.")
		autoscaling_disableMetricsCollectionCmd.MarkFlagRequired("auto-scaling-group-name")
	})
	autoscalingCmd.AddCommand(autoscaling_disableMetricsCollectionCmd)
}
