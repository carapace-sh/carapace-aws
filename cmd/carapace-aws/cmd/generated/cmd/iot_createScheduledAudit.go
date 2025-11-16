package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createScheduledAuditCmd = &cobra.Command{
	Use:   "create-scheduled-audit",
	Short: "Creates a scheduled audit that is run at a specified time interval.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createScheduledAuditCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createScheduledAuditCmd).Standalone()

		iot_createScheduledAuditCmd.Flags().String("day-of-month", "", "The day of the month on which the scheduled audit takes place.")
		iot_createScheduledAuditCmd.Flags().String("day-of-week", "", "The day of the week on which the scheduled audit takes place, either `SUN`, `MON`, `TUE`, `WED`, `THU`, `FRI`, or `SAT`.")
		iot_createScheduledAuditCmd.Flags().String("frequency", "", "How often the scheduled audit takes place, either `DAILY`, `WEEKLY`, `BIWEEKLY` or `MONTHLY`.")
		iot_createScheduledAuditCmd.Flags().String("scheduled-audit-name", "", "The name you want to give to the scheduled audit.")
		iot_createScheduledAuditCmd.Flags().String("tags", "", "Metadata that can be used to manage the scheduled audit.")
		iot_createScheduledAuditCmd.Flags().String("target-check-names", "", "Which checks are performed during the scheduled audit.")
		iot_createScheduledAuditCmd.MarkFlagRequired("frequency")
		iot_createScheduledAuditCmd.MarkFlagRequired("scheduled-audit-name")
		iot_createScheduledAuditCmd.MarkFlagRequired("target-check-names")
	})
	iotCmd.AddCommand(iot_createScheduledAuditCmd)
}
