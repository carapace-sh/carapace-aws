package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_putMetricAlarmCmd = &cobra.Command{
	Use:   "put-metric-alarm",
	Short: "Creates or updates an alarm and associates it with the specified metric, metric math expression, anomaly detection model, or Metrics Insights query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_putMetricAlarmCmd).Standalone()

	cloudwatch_putMetricAlarmCmd.Flags().String("actions-enabled", "", "Indicates whether actions should be executed during any changes to the alarm state.")
	cloudwatch_putMetricAlarmCmd.Flags().String("alarm-actions", "", "The actions to execute when this alarm transitions to the `ALARM` state from any other state.")
	cloudwatch_putMetricAlarmCmd.Flags().String("alarm-description", "", "The description for the alarm.")
	cloudwatch_putMetricAlarmCmd.Flags().String("alarm-name", "", "The name for the alarm.")
	cloudwatch_putMetricAlarmCmd.Flags().String("comparison-operator", "", "The arithmetic operation to use when comparing the specified statistic and threshold.")
	cloudwatch_putMetricAlarmCmd.Flags().String("datapoints-to-alarm", "", "The number of data points that must be breaching to trigger the alarm.")
	cloudwatch_putMetricAlarmCmd.Flags().String("dimensions", "", "The dimensions for the metric specified in `MetricName`.")
	cloudwatch_putMetricAlarmCmd.Flags().String("evaluate-low-sample-count-percentile", "", "Used only for alarms based on percentiles.")
	cloudwatch_putMetricAlarmCmd.Flags().String("evaluation-periods", "", "The number of periods over which data is compared to the specified threshold.")
	cloudwatch_putMetricAlarmCmd.Flags().String("extended-statistic", "", "The extended statistic for the metric specified in `MetricName`.")
	cloudwatch_putMetricAlarmCmd.Flags().String("insufficient-data-actions", "", "The actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state.")
	cloudwatch_putMetricAlarmCmd.Flags().String("metric-name", "", "The name for the metric associated with the alarm.")
	cloudwatch_putMetricAlarmCmd.Flags().String("metrics", "", "An array of `MetricDataQuery` structures that enable you to create an alarm based on the result of a metric math expression.")
	cloudwatch_putMetricAlarmCmd.Flags().String("namespace", "", "The namespace for the metric associated specified in `MetricName`.")
	cloudwatch_putMetricAlarmCmd.Flags().String("okactions", "", "The actions to execute when this alarm transitions to an `OK` state from any other state.")
	cloudwatch_putMetricAlarmCmd.Flags().String("period", "", "The length, in seconds, used each time the metric specified in `MetricName` is evaluated.")
	cloudwatch_putMetricAlarmCmd.Flags().String("statistic", "", "The statistic for the metric specified in `MetricName`, other than percentile.")
	cloudwatch_putMetricAlarmCmd.Flags().String("tags", "", "A list of key-value pairs to associate with the alarm.")
	cloudwatch_putMetricAlarmCmd.Flags().String("threshold", "", "The value against which the specified statistic is compared.")
	cloudwatch_putMetricAlarmCmd.Flags().String("threshold-metric-id", "", "If this is an alarm based on an anomaly detection model, make this value match the ID of the `ANOMALY_DETECTION_BAND` function.")
	cloudwatch_putMetricAlarmCmd.Flags().String("treat-missing-data", "", "Sets how this alarm is to handle missing data points.")
	cloudwatch_putMetricAlarmCmd.Flags().String("unit", "", "The unit of measure for the statistic.")
	cloudwatch_putMetricAlarmCmd.MarkFlagRequired("alarm-name")
	cloudwatch_putMetricAlarmCmd.MarkFlagRequired("comparison-operator")
	cloudwatch_putMetricAlarmCmd.MarkFlagRequired("evaluation-periods")
	cloudwatchCmd.AddCommand(cloudwatch_putMetricAlarmCmd)
}
