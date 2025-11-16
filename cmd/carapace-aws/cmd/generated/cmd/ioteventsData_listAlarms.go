package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioteventsData_listAlarmsCmd = &cobra.Command{
	Use:   "list-alarms",
	Short: "Lists one or more alarms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioteventsData_listAlarmsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ioteventsData_listAlarmsCmd).Standalone()

		ioteventsData_listAlarmsCmd.Flags().String("alarm-model-name", "", "The name of the alarm model.")
		ioteventsData_listAlarmsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		ioteventsData_listAlarmsCmd.Flags().String("next-token", "", "The token that you can use to return the next set of results.")
		ioteventsData_listAlarmsCmd.MarkFlagRequired("alarm-model-name")
	})
	ioteventsDataCmd.AddCommand(ioteventsData_listAlarmsCmd)
}
