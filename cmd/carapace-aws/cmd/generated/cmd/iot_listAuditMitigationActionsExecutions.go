package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listAuditMitigationActionsExecutionsCmd = &cobra.Command{
	Use:   "list-audit-mitigation-actions-executions",
	Short: "Gets the status of audit mitigation action tasks that were executed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listAuditMitigationActionsExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listAuditMitigationActionsExecutionsCmd).Standalone()

		iot_listAuditMitigationActionsExecutionsCmd.Flags().String("action-status", "", "Specify this filter to limit results to those with a specific status.")
		iot_listAuditMitigationActionsExecutionsCmd.Flags().String("finding-id", "", "Specify this filter to limit results to those that were applied to a specific audit finding.")
		iot_listAuditMitigationActionsExecutionsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iot_listAuditMitigationActionsExecutionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		iot_listAuditMitigationActionsExecutionsCmd.Flags().String("task-id", "", "Specify this filter to limit results to actions for a specific audit mitigation actions task.")
		iot_listAuditMitigationActionsExecutionsCmd.MarkFlagRequired("finding-id")
		iot_listAuditMitigationActionsExecutionsCmd.MarkFlagRequired("task-id")
	})
	iotCmd.AddCommand(iot_listAuditMitigationActionsExecutionsCmd)
}
