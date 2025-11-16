package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_startAuditMitigationActionsTaskCmd = &cobra.Command{
	Use:   "start-audit-mitigation-actions-task",
	Short: "Starts a task that applies a set of mitigation actions to the specified target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_startAuditMitigationActionsTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_startAuditMitigationActionsTaskCmd).Standalone()

		iot_startAuditMitigationActionsTaskCmd.Flags().String("audit-check-to-actions-mapping", "", "For an audit check, specifies which mitigation actions to apply.")
		iot_startAuditMitigationActionsTaskCmd.Flags().String("client-request-token", "", "Each audit mitigation task must have a unique client request token.")
		iot_startAuditMitigationActionsTaskCmd.Flags().String("target", "", "Specifies the audit findings to which the mitigation actions are applied.")
		iot_startAuditMitigationActionsTaskCmd.Flags().String("task-id", "", "A unique identifier for the task.")
		iot_startAuditMitigationActionsTaskCmd.MarkFlagRequired("audit-check-to-actions-mapping")
		iot_startAuditMitigationActionsTaskCmd.MarkFlagRequired("client-request-token")
		iot_startAuditMitigationActionsTaskCmd.MarkFlagRequired("target")
		iot_startAuditMitigationActionsTaskCmd.MarkFlagRequired("task-id")
	})
	iotCmd.AddCommand(iot_startAuditMitigationActionsTaskCmd)
}
