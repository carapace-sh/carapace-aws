package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_cancelAuditMitigationActionsTaskCmd = &cobra.Command{
	Use:   "cancel-audit-mitigation-actions-task",
	Short: "Cancels a mitigation action task that is in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_cancelAuditMitigationActionsTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_cancelAuditMitigationActionsTaskCmd).Standalone()

		iot_cancelAuditMitigationActionsTaskCmd.Flags().String("task-id", "", "The unique identifier for the task that you want to cancel.")
		iot_cancelAuditMitigationActionsTaskCmd.MarkFlagRequired("task-id")
	})
	iotCmd.AddCommand(iot_cancelAuditMitigationActionsTaskCmd)
}
