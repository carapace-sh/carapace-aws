package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_deleteAlarmModelCmd = &cobra.Command{
	Use:   "delete-alarm-model",
	Short: "Deletes an alarm model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_deleteAlarmModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotevents_deleteAlarmModelCmd).Standalone()

		iotevents_deleteAlarmModelCmd.Flags().String("alarm-model-name", "", "The name of the alarm model.")
		iotevents_deleteAlarmModelCmd.MarkFlagRequired("alarm-model-name")
	})
	ioteventsCmd.AddCommand(iotevents_deleteAlarmModelCmd)
}
