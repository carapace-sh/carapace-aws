package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteScheduledAuditCmd = &cobra.Command{
	Use:   "delete-scheduled-audit",
	Short: "Deletes a scheduled audit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteScheduledAuditCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteScheduledAuditCmd).Standalone()

		iot_deleteScheduledAuditCmd.Flags().String("scheduled-audit-name", "", "The name of the scheduled audit you want to delete.")
		iot_deleteScheduledAuditCmd.MarkFlagRequired("scheduled-audit-name")
	})
	iotCmd.AddCommand(iot_deleteScheduledAuditCmd)
}
