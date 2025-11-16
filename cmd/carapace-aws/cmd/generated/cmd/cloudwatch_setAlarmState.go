package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_setAlarmStateCmd = &cobra.Command{
	Use:   "set-alarm-state",
	Short: "Temporarily sets the state of an alarm for testing purposes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_setAlarmStateCmd).Standalone()

	cloudwatch_setAlarmStateCmd.Flags().String("alarm-name", "", "The name of the alarm.")
	cloudwatch_setAlarmStateCmd.Flags().String("state-reason", "", "The reason that this alarm is set to this specific state, in text format.")
	cloudwatch_setAlarmStateCmd.Flags().String("state-reason-data", "", "The reason that this alarm is set to this specific state, in JSON format.")
	cloudwatch_setAlarmStateCmd.Flags().String("state-value", "", "The value of the state.")
	cloudwatch_setAlarmStateCmd.MarkFlagRequired("alarm-name")
	cloudwatch_setAlarmStateCmd.MarkFlagRequired("state-reason")
	cloudwatch_setAlarmStateCmd.MarkFlagRequired("state-value")
	cloudwatchCmd.AddCommand(cloudwatch_setAlarmStateCmd)
}
