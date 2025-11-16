package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteReplicationTaskCmd = &cobra.Command{
	Use:   "delete-replication-task",
	Short: "Deletes the specified replication task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteReplicationTaskCmd).Standalone()

	dms_deleteReplicationTaskCmd.Flags().String("replication-task-arn", "", "The Amazon Resource Name (ARN) of the replication task to be deleted.")
	dms_deleteReplicationTaskCmd.MarkFlagRequired("replication-task-arn")
	dmsCmd.AddCommand(dms_deleteReplicationTaskCmd)
}
