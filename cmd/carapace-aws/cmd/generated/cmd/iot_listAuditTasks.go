package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listAuditTasksCmd = &cobra.Command{
	Use:   "list-audit-tasks",
	Short: "Lists the Device Defender audits that have been performed during a given time period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listAuditTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listAuditTasksCmd).Standalone()

		iot_listAuditTasksCmd.Flags().String("end-time", "", "The end of the time period.")
		iot_listAuditTasksCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iot_listAuditTasksCmd.Flags().String("next-token", "", "The token for the next set of results.")
		iot_listAuditTasksCmd.Flags().String("start-time", "", "The beginning of the time period.")
		iot_listAuditTasksCmd.Flags().String("task-status", "", "A filter to limit the output to audits with the specified completion status: can be one of \"IN\\_PROGRESS\", \"COMPLETED\", \"FAILED\", or \"CANCELED\".")
		iot_listAuditTasksCmd.Flags().String("task-type", "", "A filter to limit the output to the specified type of audit: can be one of \"ON\\_DEMAND\\_AUDIT\\_TASK\" or \"SCHEDULED\\_\\_AUDIT\\_TASK\".")
		iot_listAuditTasksCmd.MarkFlagRequired("end-time")
		iot_listAuditTasksCmd.MarkFlagRequired("start-time")
	})
	iotCmd.AddCommand(iot_listAuditTasksCmd)
}
