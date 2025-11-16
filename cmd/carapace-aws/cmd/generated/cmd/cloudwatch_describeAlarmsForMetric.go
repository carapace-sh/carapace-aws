package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_describeAlarmsForMetricCmd = &cobra.Command{
	Use:   "describe-alarms-for-metric",
	Short: "Retrieves the alarms for the specified metric.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_describeAlarmsForMetricCmd).Standalone()

	cloudwatch_describeAlarmsForMetricCmd.Flags().String("dimensions", "", "The dimensions associated with the metric.")
	cloudwatch_describeAlarmsForMetricCmd.Flags().String("extended-statistic", "", "The percentile statistic for the metric.")
	cloudwatch_describeAlarmsForMetricCmd.Flags().String("metric-name", "", "The name of the metric.")
	cloudwatch_describeAlarmsForMetricCmd.Flags().String("namespace", "", "The namespace of the metric.")
	cloudwatch_describeAlarmsForMetricCmd.Flags().String("period", "", "The period, in seconds, over which the statistic is applied.")
	cloudwatch_describeAlarmsForMetricCmd.Flags().String("statistic", "", "The statistic for the metric, other than percentiles.")
	cloudwatch_describeAlarmsForMetricCmd.Flags().String("unit", "", "The unit for the metric.")
	cloudwatch_describeAlarmsForMetricCmd.MarkFlagRequired("metric-name")
	cloudwatch_describeAlarmsForMetricCmd.MarkFlagRequired("namespace")
	cloudwatchCmd.AddCommand(cloudwatch_describeAlarmsForMetricCmd)
}
