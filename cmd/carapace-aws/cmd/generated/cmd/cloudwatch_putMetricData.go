package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_putMetricDataCmd = &cobra.Command{
	Use:   "put-metric-data",
	Short: "Publishes metric data to Amazon CloudWatch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_putMetricDataCmd).Standalone()

	cloudwatch_putMetricDataCmd.Flags().String("entity-metric-data", "", "Data for metrics that contain associated entity information.")
	cloudwatch_putMetricDataCmd.Flags().String("metric-data", "", "The data for the metrics.")
	cloudwatch_putMetricDataCmd.Flags().String("namespace", "", "The namespace for the metric data.")
	cloudwatch_putMetricDataCmd.Flags().String("strict-entity-validation", "", "Whether to accept valid metric data when an invalid entity is sent.")
	cloudwatch_putMetricDataCmd.MarkFlagRequired("namespace")
	cloudwatchCmd.AddCommand(cloudwatch_putMetricDataCmd)
}
