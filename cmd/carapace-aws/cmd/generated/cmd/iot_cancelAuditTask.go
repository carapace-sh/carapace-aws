package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_cancelAuditTaskCmd = &cobra.Command{
	Use:   "cancel-audit-task",
	Short: "Cancels an audit that is in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_cancelAuditTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_cancelAuditTaskCmd).Standalone()

		iot_cancelAuditTaskCmd.Flags().String("task-id", "", "The ID of the audit you want to cancel.")
		iot_cancelAuditTaskCmd.MarkFlagRequired("task-id")
	})
	iotCmd.AddCommand(iot_cancelAuditTaskCmd)
}
