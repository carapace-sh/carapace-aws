package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeReplicationTaskAssessmentRunsCmd = &cobra.Command{
	Use:   "describe-replication-task-assessment-runs",
	Short: "Returns a paginated list of premigration assessment runs based on filter settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeReplicationTaskAssessmentRunsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeReplicationTaskAssessmentRunsCmd).Standalone()

		dms_describeReplicationTaskAssessmentRunsCmd.Flags().String("filters", "", "Filters applied to the premigration assessment runs described in the form of key-value pairs.")
		dms_describeReplicationTaskAssessmentRunsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		dms_describeReplicationTaskAssessmentRunsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	dmsCmd.AddCommand(dms_describeReplicationTaskAssessmentRunsCmd)
}
