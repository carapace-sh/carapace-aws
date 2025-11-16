package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_startReplicationTaskAssessmentCmd = &cobra.Command{
	Use:   "start-replication-task-assessment",
	Short: "Starts the replication task assessment for unsupported data types in the source database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_startReplicationTaskAssessmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_startReplicationTaskAssessmentCmd).Standalone()

		dms_startReplicationTaskAssessmentCmd.Flags().String("replication-task-arn", "", "The Amazon Resource Name (ARN) of the replication task.")
		dms_startReplicationTaskAssessmentCmd.MarkFlagRequired("replication-task-arn")
	})
	dmsCmd.AddCommand(dms_startReplicationTaskAssessmentCmd)
}
