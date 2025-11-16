package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioteventsData_batchSnoozeAlarmCmd = &cobra.Command{
	Use:   "batch-snooze-alarm",
	Short: "Changes one or more alarms to the snooze mode.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioteventsData_batchSnoozeAlarmCmd).Standalone()

	ioteventsData_batchSnoozeAlarmCmd.Flags().String("snooze-action-requests", "", "The list of snooze action requests.")
	ioteventsData_batchSnoozeAlarmCmd.MarkFlagRequired("snooze-action-requests")
	ioteventsDataCmd.AddCommand(ioteventsData_batchSnoozeAlarmCmd)
}
