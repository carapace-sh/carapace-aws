package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_createAlarmModelCmd = &cobra.Command{
	Use:   "create-alarm-model",
	Short: "Creates an alarm model to monitor an AWS IoT Events input attribute.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_createAlarmModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotevents_createAlarmModelCmd).Standalone()

		iotevents_createAlarmModelCmd.Flags().String("alarm-capabilities", "", "Contains the configuration information of alarm state changes.")
		iotevents_createAlarmModelCmd.Flags().String("alarm-event-actions", "", "Contains information about one or more alarm actions.")
		iotevents_createAlarmModelCmd.Flags().String("alarm-model-description", "", "A description that tells you what the alarm model detects.")
		iotevents_createAlarmModelCmd.Flags().String("alarm-model-name", "", "A unique name that helps you identify the alarm model.")
		iotevents_createAlarmModelCmd.Flags().String("alarm-notification", "", "Contains information about one or more notification actions.")
		iotevents_createAlarmModelCmd.Flags().String("alarm-rule", "", "Defines when your alarm is invoked.")
		iotevents_createAlarmModelCmd.Flags().String("key", "", "An input attribute used as a key to create an alarm.")
		iotevents_createAlarmModelCmd.Flags().String("role-arn", "", "The ARN of the IAM role that allows the alarm to perform actions and access AWS resources.")
		iotevents_createAlarmModelCmd.Flags().String("severity", "", "A non-negative integer that reflects the severity level of the alarm.")
		iotevents_createAlarmModelCmd.Flags().String("tags", "", "A list of key-value pairs that contain metadata for the alarm model.")
		iotevents_createAlarmModelCmd.MarkFlagRequired("alarm-model-name")
		iotevents_createAlarmModelCmd.MarkFlagRequired("alarm-rule")
		iotevents_createAlarmModelCmd.MarkFlagRequired("role-arn")
	})
	ioteventsCmd.AddCommand(iotevents_createAlarmModelCmd)
}
