package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioteventsData_batchEnableAlarmCmd = &cobra.Command{
	Use:   "batch-enable-alarm",
	Short: "Enables one or more alarms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioteventsData_batchEnableAlarmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ioteventsData_batchEnableAlarmCmd).Standalone()

		ioteventsData_batchEnableAlarmCmd.Flags().String("enable-action-requests", "", "The list of enable action requests.")
		ioteventsData_batchEnableAlarmCmd.MarkFlagRequired("enable-action-requests")
	})
	ioteventsDataCmd.AddCommand(ioteventsData_batchEnableAlarmCmd)
}
