package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioteventsData_batchResetAlarmCmd = &cobra.Command{
	Use:   "batch-reset-alarm",
	Short: "Resets one or more alarms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioteventsData_batchResetAlarmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ioteventsData_batchResetAlarmCmd).Standalone()

		ioteventsData_batchResetAlarmCmd.Flags().String("reset-action-requests", "", "The list of reset action requests.")
		ioteventsData_batchResetAlarmCmd.MarkFlagRequired("reset-action-requests")
	})
	ioteventsDataCmd.AddCommand(ioteventsData_batchResetAlarmCmd)
}
