package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteReplicationTaskAssessmentRunCmd = &cobra.Command{
	Use:   "delete-replication-task-assessment-run",
	Short: "Deletes the record of a single premigration assessment run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteReplicationTaskAssessmentRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_deleteReplicationTaskAssessmentRunCmd).Standalone()

		dms_deleteReplicationTaskAssessmentRunCmd.Flags().String("replication-task-assessment-run-arn", "", "Amazon Resource Name (ARN) of the premigration assessment run to be deleted.")
		dms_deleteReplicationTaskAssessmentRunCmd.MarkFlagRequired("replication-task-assessment-run-arn")
	})
	dmsCmd.AddCommand(dms_deleteReplicationTaskAssessmentRunCmd)
}
