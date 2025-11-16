package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioteventsData_batchDisableAlarmCmd = &cobra.Command{
	Use:   "batch-disable-alarm",
	Short: "Disables one or more alarms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioteventsData_batchDisableAlarmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ioteventsData_batchDisableAlarmCmd).Standalone()

		ioteventsData_batchDisableAlarmCmd.Flags().String("disable-action-requests", "", "The list of disable action requests.")
		ioteventsData_batchDisableAlarmCmd.MarkFlagRequired("disable-action-requests")
	})
	ioteventsDataCmd.AddCommand(ioteventsData_batchDisableAlarmCmd)
}
