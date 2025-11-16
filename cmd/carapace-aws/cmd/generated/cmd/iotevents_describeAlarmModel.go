package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_describeAlarmModelCmd = &cobra.Command{
	Use:   "describe-alarm-model",
	Short: "Retrieves information about an alarm model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_describeAlarmModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotevents_describeAlarmModelCmd).Standalone()

		iotevents_describeAlarmModelCmd.Flags().String("alarm-model-name", "", "The name of the alarm model.")
		iotevents_describeAlarmModelCmd.Flags().String("alarm-model-version", "", "The version of the alarm model.")
		iotevents_describeAlarmModelCmd.MarkFlagRequired("alarm-model-name")
	})
	ioteventsCmd.AddCommand(iotevents_describeAlarmModelCmd)
}
