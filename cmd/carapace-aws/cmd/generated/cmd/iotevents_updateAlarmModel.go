package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_updateAlarmModelCmd = &cobra.Command{
	Use:   "update-alarm-model",
	Short: "Updates an alarm model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_updateAlarmModelCmd).Standalone()

	iotevents_updateAlarmModelCmd.Flags().String("alarm-capabilities", "", "Contains the configuration information of alarm state changes.")
	iotevents_updateAlarmModelCmd.Flags().String("alarm-event-actions", "", "Contains information about one or more alarm actions.")
	iotevents_updateAlarmModelCmd.Flags().String("alarm-model-description", "", "The description of the alarm model.")
	iotevents_updateAlarmModelCmd.Flags().String("alarm-model-name", "", "The name of the alarm model.")
	iotevents_updateAlarmModelCmd.Flags().String("alarm-notification", "", "Contains information about one or more notification actions.")
	iotevents_updateAlarmModelCmd.Flags().String("alarm-rule", "", "Defines when your alarm is invoked.")
	iotevents_updateAlarmModelCmd.Flags().String("role-arn", "", "The ARN of the IAM role that allows the alarm to perform actions and access AWS resources.")
	iotevents_updateAlarmModelCmd.Flags().String("severity", "", "A non-negative integer that reflects the severity level of the alarm.")
	iotevents_updateAlarmModelCmd.MarkFlagRequired("alarm-model-name")
	iotevents_updateAlarmModelCmd.MarkFlagRequired("alarm-rule")
	iotevents_updateAlarmModelCmd.MarkFlagRequired("role-arn")
	ioteventsCmd.AddCommand(iotevents_updateAlarmModelCmd)
}
