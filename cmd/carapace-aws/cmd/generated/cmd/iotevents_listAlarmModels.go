package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_listAlarmModelsCmd = &cobra.Command{
	Use:   "list-alarm-models",
	Short: "Lists the alarm models that you created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_listAlarmModelsCmd).Standalone()

	iotevents_listAlarmModelsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	iotevents_listAlarmModelsCmd.Flags().String("next-token", "", "The token that you can use to return the next set of results.")
	ioteventsCmd.AddCommand(iotevents_listAlarmModelsCmd)
}
