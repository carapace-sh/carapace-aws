package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeReplicationInstanceTaskLogsCmd = &cobra.Command{
	Use:   "describe-replication-instance-task-logs",
	Short: "Returns information about the task logs for the specified task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeReplicationInstanceTaskLogsCmd).Standalone()

	dms_describeReplicationInstanceTaskLogsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	dms_describeReplicationInstanceTaskLogsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dms_describeReplicationInstanceTaskLogsCmd.Flags().String("replication-instance-arn", "", "The Amazon Resource Name (ARN) of the replication instance.")
	dms_describeReplicationInstanceTaskLogsCmd.MarkFlagRequired("replication-instance-arn")
	dmsCmd.AddCommand(dms_describeReplicationInstanceTaskLogsCmd)
}
