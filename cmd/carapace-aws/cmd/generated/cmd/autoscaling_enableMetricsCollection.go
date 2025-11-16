package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_enableMetricsCollectionCmd = &cobra.Command{
	Use:   "enable-metrics-collection",
	Short: "Enables group metrics collection for the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_enableMetricsCollectionCmd).Standalone()

	autoscaling_enableMetricsCollectionCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_enableMetricsCollectionCmd.Flags().String("granularity", "", "The frequency at which Amazon EC2 Auto Scaling sends aggregated data to CloudWatch.")
	autoscaling_enableMetricsCollectionCmd.Flags().String("metrics", "", "Identifies the metrics to enable.")
	autoscaling_enableMetricsCollectionCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscaling_enableMetricsCollectionCmd.MarkFlagRequired("granularity")
	autoscalingCmd.AddCommand(autoscaling_enableMetricsCollectionCmd)
}
