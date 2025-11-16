package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_moveReplicationTaskCmd = &cobra.Command{
	Use:   "move-replication-task",
	Short: "Moves a replication task from its current replication instance to a different target replication instance using the specified parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_moveReplicationTaskCmd).Standalone()

	dms_moveReplicationTaskCmd.Flags().String("replication-task-arn", "", "The Amazon Resource Name (ARN) of the task that you want to move.")
	dms_moveReplicationTaskCmd.Flags().String("target-replication-instance-arn", "", "The ARN of the replication instance where you want to move the task to.")
	dms_moveReplicationTaskCmd.MarkFlagRequired("replication-task-arn")
	dms_moveReplicationTaskCmd.MarkFlagRequired("target-replication-instance-arn")
	dmsCmd.AddCommand(dms_moveReplicationTaskCmd)
}
