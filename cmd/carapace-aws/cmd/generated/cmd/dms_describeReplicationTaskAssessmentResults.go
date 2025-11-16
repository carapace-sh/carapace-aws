package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeReplicationTaskAssessmentResultsCmd = &cobra.Command{
	Use:   "describe-replication-task-assessment-results",
	Short: "Returns the task assessment results from the Amazon S3 bucket that DMS creates in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeReplicationTaskAssessmentResultsCmd).Standalone()

	dms_describeReplicationTaskAssessmentResultsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	dms_describeReplicationTaskAssessmentResultsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dms_describeReplicationTaskAssessmentResultsCmd.Flags().String("replication-task-arn", "", "The Amazon Resource Name (ARN) string that uniquely identifies the task.")
	dmsCmd.AddCommand(dms_describeReplicationTaskAssessmentResultsCmd)
}
