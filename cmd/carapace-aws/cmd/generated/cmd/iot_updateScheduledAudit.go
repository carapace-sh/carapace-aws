package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateScheduledAuditCmd = &cobra.Command{
	Use:   "update-scheduled-audit",
	Short: "Updates a scheduled audit, including which checks are performed and how often the audit takes place.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateScheduledAuditCmd).Standalone()

	iot_updateScheduledAuditCmd.Flags().String("day-of-month", "", "The day of the month on which the scheduled audit takes place.")
	iot_updateScheduledAuditCmd.Flags().String("day-of-week", "", "The day of the week on which the scheduled audit takes place.")
	iot_updateScheduledAuditCmd.Flags().String("frequency", "", "How often the scheduled audit takes place, either `DAILY`, `WEEKLY`, `BIWEEKLY`, or `MONTHLY`.")
	iot_updateScheduledAuditCmd.Flags().String("scheduled-audit-name", "", "The name of the scheduled audit.")
	iot_updateScheduledAuditCmd.Flags().String("target-check-names", "", "Which checks are performed during the scheduled audit.")
	iot_updateScheduledAuditCmd.MarkFlagRequired("scheduled-audit-name")
	iotCmd.AddCommand(iot_updateScheduledAuditCmd)
}
