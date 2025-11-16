package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listAuditMitigationActionsTasksCmd = &cobra.Command{
	Use:   "list-audit-mitigation-actions-tasks",
	Short: "Gets a list of audit mitigation action tasks that match the specified filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listAuditMitigationActionsTasksCmd).Standalone()

	iot_listAuditMitigationActionsTasksCmd.Flags().String("audit-task-id", "", "Specify this filter to limit results to tasks that were applied to results for a specific audit.")
	iot_listAuditMitigationActionsTasksCmd.Flags().String("end-time", "", "Specify this filter to limit results to tasks that were completed or canceled on or before a specific date and time.")
	iot_listAuditMitigationActionsTasksCmd.Flags().String("finding-id", "", "Specify this filter to limit results to tasks that were applied to a specific audit finding.")
	iot_listAuditMitigationActionsTasksCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listAuditMitigationActionsTasksCmd.Flags().String("next-token", "", "The token for the next set of results.")
	iot_listAuditMitigationActionsTasksCmd.Flags().String("start-time", "", "Specify this filter to limit results to tasks that began on or after a specific date and time.")
	iot_listAuditMitigationActionsTasksCmd.Flags().String("task-status", "", "Specify this filter to limit results to tasks that are in a specific state.")
	iot_listAuditMitigationActionsTasksCmd.MarkFlagRequired("end-time")
	iot_listAuditMitigationActionsTasksCmd.MarkFlagRequired("start-time")
	iotCmd.AddCommand(iot_listAuditMitigationActionsTasksCmd)
}
