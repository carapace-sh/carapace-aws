package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_listAlarmModelVersionsCmd = &cobra.Command{
	Use:   "list-alarm-model-versions",
	Short: "Lists all the versions of an alarm model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_listAlarmModelVersionsCmd).Standalone()

	iotevents_listAlarmModelVersionsCmd.Flags().String("alarm-model-name", "", "The name of the alarm model.")
	iotevents_listAlarmModelVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	iotevents_listAlarmModelVersionsCmd.Flags().String("next-token", "", "The token that you can use to return the next set of results.")
	iotevents_listAlarmModelVersionsCmd.MarkFlagRequired("alarm-model-name")
	ioteventsCmd.AddCommand(iotevents_listAlarmModelVersionsCmd)
}
