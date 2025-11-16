package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_startReplicationTaskCmd = &cobra.Command{
	Use:   "start-replication-task",
	Short: "Starts the replication task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_startReplicationTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_startReplicationTaskCmd).Standalone()

		dms_startReplicationTaskCmd.Flags().String("cdc-start-position", "", "Indicates when you want a change data capture (CDC) operation to start.")
		dms_startReplicationTaskCmd.Flags().String("cdc-start-time", "", "Indicates the start time for a change data capture (CDC) operation.")
		dms_startReplicationTaskCmd.Flags().String("cdc-stop-position", "", "Indicates when you want a change data capture (CDC) operation to stop.")
		dms_startReplicationTaskCmd.Flags().String("replication-task-arn", "", "The Amazon Resource Name (ARN) of the replication task to be started.")
		dms_startReplicationTaskCmd.Flags().String("start-replication-task-type", "", "The type of replication task to start.")
		dms_startReplicationTaskCmd.MarkFlagRequired("replication-task-arn")
		dms_startReplicationTaskCmd.MarkFlagRequired("start-replication-task-type")
	})
	dmsCmd.AddCommand(dms_startReplicationTaskCmd)
}
