package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_createReplicationTaskCmd = &cobra.Command{
	Use:   "create-replication-task",
	Short: "Creates a replication task using the specified parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_createReplicationTaskCmd).Standalone()

	dms_createReplicationTaskCmd.Flags().String("cdc-start-position", "", "Indicates when you want a change data capture (CDC) operation to start.")
	dms_createReplicationTaskCmd.Flags().String("cdc-start-time", "", "Indicates the start time for a change data capture (CDC) operation.")
	dms_createReplicationTaskCmd.Flags().String("cdc-stop-position", "", "Indicates when you want a change data capture (CDC) operation to stop.")
	dms_createReplicationTaskCmd.Flags().String("migration-type", "", "The migration type.")
	dms_createReplicationTaskCmd.Flags().String("replication-instance-arn", "", "The Amazon Resource Name (ARN) of a replication instance.")
	dms_createReplicationTaskCmd.Flags().String("replication-task-identifier", "", "An identifier for the replication task.")
	dms_createReplicationTaskCmd.Flags().String("replication-task-settings", "", "Overall settings for the task, in JSON format.")
	dms_createReplicationTaskCmd.Flags().String("resource-identifier", "", "A friendly name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.")
	dms_createReplicationTaskCmd.Flags().String("source-endpoint-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies the source endpoint.")
	dms_createReplicationTaskCmd.Flags().String("table-mappings", "", "The table mappings for the task, in JSON format.")
	dms_createReplicationTaskCmd.Flags().String("tags", "", "One or more tags to be assigned to the replication task.")
	dms_createReplicationTaskCmd.Flags().String("target-endpoint-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies the target endpoint.")
	dms_createReplicationTaskCmd.Flags().String("task-data", "", "Supplemental information that the task requires to migrate the data for certain source and target endpoints.")
	dms_createReplicationTaskCmd.MarkFlagRequired("migration-type")
	dms_createReplicationTaskCmd.MarkFlagRequired("replication-instance-arn")
	dms_createReplicationTaskCmd.MarkFlagRequired("replication-task-identifier")
	dms_createReplicationTaskCmd.MarkFlagRequired("source-endpoint-arn")
	dms_createReplicationTaskCmd.MarkFlagRequired("table-mappings")
	dms_createReplicationTaskCmd.MarkFlagRequired("target-endpoint-arn")
	dmsCmd.AddCommand(dms_createReplicationTaskCmd)
}
