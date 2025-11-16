package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_modifyReplicationTaskCmd = &cobra.Command{
	Use:   "modify-replication-task",
	Short: "Modifies the specified replication task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_modifyReplicationTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_modifyReplicationTaskCmd).Standalone()

		dms_modifyReplicationTaskCmd.Flags().String("cdc-start-position", "", "Indicates when you want a change data capture (CDC) operation to start.")
		dms_modifyReplicationTaskCmd.Flags().String("cdc-start-time", "", "Indicates the start time for a change data capture (CDC) operation.")
		dms_modifyReplicationTaskCmd.Flags().String("cdc-stop-position", "", "Indicates when you want a change data capture (CDC) operation to stop.")
		dms_modifyReplicationTaskCmd.Flags().String("migration-type", "", "The migration type.")
		dms_modifyReplicationTaskCmd.Flags().String("replication-task-arn", "", "The Amazon Resource Name (ARN) of the replication task.")
		dms_modifyReplicationTaskCmd.Flags().String("replication-task-identifier", "", "The replication task identifier.")
		dms_modifyReplicationTaskCmd.Flags().String("replication-task-settings", "", "JSON file that contains settings for the task, such as task metadata settings.")
		dms_modifyReplicationTaskCmd.Flags().String("table-mappings", "", "When using the CLI or boto3, provide the path of the JSON file that contains the table mappings.")
		dms_modifyReplicationTaskCmd.Flags().String("task-data", "", "Supplemental information that the task requires to migrate the data for certain source and target endpoints.")
		dms_modifyReplicationTaskCmd.MarkFlagRequired("replication-task-arn")
	})
	dmsCmd.AddCommand(dms_modifyReplicationTaskCmd)
}
