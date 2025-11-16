package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeAuditMitigationActionsTaskCmd = &cobra.Command{
	Use:   "describe-audit-mitigation-actions-task",
	Short: "Gets information about an audit mitigation task that is used to apply mitigation actions to a set of audit findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeAuditMitigationActionsTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeAuditMitigationActionsTaskCmd).Standalone()

		iot_describeAuditMitigationActionsTaskCmd.Flags().String("task-id", "", "The unique identifier for the audit mitigation task.")
		iot_describeAuditMitigationActionsTaskCmd.MarkFlagRequired("task-id")
	})
	iotCmd.AddCommand(iot_describeAuditMitigationActionsTaskCmd)
}
