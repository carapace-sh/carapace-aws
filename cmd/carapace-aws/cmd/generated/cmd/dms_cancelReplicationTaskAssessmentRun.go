package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_cancelReplicationTaskAssessmentRunCmd = &cobra.Command{
	Use:   "cancel-replication-task-assessment-run",
	Short: "Cancels a single premigration assessment run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_cancelReplicationTaskAssessmentRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_cancelReplicationTaskAssessmentRunCmd).Standalone()

		dms_cancelReplicationTaskAssessmentRunCmd.Flags().String("replication-task-assessment-run-arn", "", "Amazon Resource Name (ARN) of the premigration assessment run to be canceled.")
		dms_cancelReplicationTaskAssessmentRunCmd.MarkFlagRequired("replication-task-assessment-run-arn")
	})
	dmsCmd.AddCommand(dms_cancelReplicationTaskAssessmentRunCmd)
}
