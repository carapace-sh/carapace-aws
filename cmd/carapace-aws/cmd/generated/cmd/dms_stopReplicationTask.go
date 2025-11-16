package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_stopReplicationTaskCmd = &cobra.Command{
	Use:   "stop-replication-task",
	Short: "Stops the replication task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_stopReplicationTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_stopReplicationTaskCmd).Standalone()

		dms_stopReplicationTaskCmd.Flags().String("replication-task-arn", "", "The Amazon Resource Name(ARN) of the replication task to be stopped.")
		dms_stopReplicationTaskCmd.MarkFlagRequired("replication-task-arn")
	})
	dmsCmd.AddCommand(dms_stopReplicationTaskCmd)
}
