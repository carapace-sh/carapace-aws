package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_describeAlarmHistoryCmd = &cobra.Command{
	Use:   "describe-alarm-history",
	Short: "Retrieves the history for the specified alarm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_describeAlarmHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_describeAlarmHistoryCmd).Standalone()

		cloudwatch_describeAlarmHistoryCmd.Flags().String("alarm-contributor-id", "", "The unique identifier of a specific alarm contributor to filter the alarm history results.")
		cloudwatch_describeAlarmHistoryCmd.Flags().String("alarm-name", "", "The name of the alarm.")
		cloudwatch_describeAlarmHistoryCmd.Flags().String("alarm-types", "", "Use this parameter to specify whether you want the operation to return metric alarms or composite alarms.")
		cloudwatch_describeAlarmHistoryCmd.Flags().String("end-date", "", "The ending date to retrieve alarm history.")
		cloudwatch_describeAlarmHistoryCmd.Flags().String("history-item-type", "", "The type of alarm histories to retrieve.")
		cloudwatch_describeAlarmHistoryCmd.Flags().String("max-records", "", "The maximum number of alarm history records to retrieve.")
		cloudwatch_describeAlarmHistoryCmd.Flags().String("next-token", "", "The token returned by a previous call to indicate that there is more data available.")
		cloudwatch_describeAlarmHistoryCmd.Flags().String("scan-by", "", "Specified whether to return the newest or oldest alarm history first.")
		cloudwatch_describeAlarmHistoryCmd.Flags().String("start-date", "", "The starting date to retrieve alarm history.")
	})
	cloudwatchCmd.AddCommand(cloudwatch_describeAlarmHistoryCmd)
}
