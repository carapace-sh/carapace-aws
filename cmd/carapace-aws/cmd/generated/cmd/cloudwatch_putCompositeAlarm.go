package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_putCompositeAlarmCmd = &cobra.Command{
	Use:   "put-composite-alarm",
	Short: "Creates or updates a *composite alarm*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_putCompositeAlarmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_putCompositeAlarmCmd).Standalone()

		cloudwatch_putCompositeAlarmCmd.Flags().String("actions-enabled", "", "Indicates whether actions should be executed during any changes to the alarm state of the composite alarm.")
		cloudwatch_putCompositeAlarmCmd.Flags().String("actions-suppressor", "", "Actions will be suppressed if the suppressor alarm is in the `ALARM` state.")
		cloudwatch_putCompositeAlarmCmd.Flags().String("actions-suppressor-extension-period", "", "The maximum time in seconds that the composite alarm waits after suppressor alarm goes out of the `ALARM` state.")
		cloudwatch_putCompositeAlarmCmd.Flags().String("actions-suppressor-wait-period", "", "The maximum time in seconds that the composite alarm waits for the suppressor alarm to go into the `ALARM` state.")
		cloudwatch_putCompositeAlarmCmd.Flags().String("alarm-actions", "", "The actions to execute when this alarm transitions to the `ALARM` state from any other state.")
		cloudwatch_putCompositeAlarmCmd.Flags().String("alarm-description", "", "The description for the composite alarm.")
		cloudwatch_putCompositeAlarmCmd.Flags().String("alarm-name", "", "The name for the composite alarm.")
		cloudwatch_putCompositeAlarmCmd.Flags().String("alarm-rule", "", "An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state.")
		cloudwatch_putCompositeAlarmCmd.Flags().String("insufficient-data-actions", "", "The actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state.")
		cloudwatch_putCompositeAlarmCmd.Flags().String("okactions", "", "The actions to execute when this alarm transitions to an `OK` state from any other state.")
		cloudwatch_putCompositeAlarmCmd.Flags().String("tags", "", "A list of key-value pairs to associate with the alarm.")
		cloudwatch_putCompositeAlarmCmd.MarkFlagRequired("alarm-name")
		cloudwatch_putCompositeAlarmCmd.MarkFlagRequired("alarm-rule")
	})
	cloudwatchCmd.AddCommand(cloudwatch_putCompositeAlarmCmd)
}
