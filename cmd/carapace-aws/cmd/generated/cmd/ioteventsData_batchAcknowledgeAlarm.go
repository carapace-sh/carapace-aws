package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioteventsData_batchAcknowledgeAlarmCmd = &cobra.Command{
	Use:   "batch-acknowledge-alarm",
	Short: "Acknowledges one or more alarms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioteventsData_batchAcknowledgeAlarmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ioteventsData_batchAcknowledgeAlarmCmd).Standalone()

		ioteventsData_batchAcknowledgeAlarmCmd.Flags().String("acknowledge-action-requests", "", "The list of acknowledge action requests.")
		ioteventsData_batchAcknowledgeAlarmCmd.MarkFlagRequired("acknowledge-action-requests")
	})
	ioteventsDataCmd.AddCommand(ioteventsData_batchAcknowledgeAlarmCmd)
}
