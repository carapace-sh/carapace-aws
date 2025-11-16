package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeReplicationTaskIndividualAssessmentsCmd = &cobra.Command{
	Use:   "describe-replication-task-individual-assessments",
	Short: "Returns a paginated list of individual assessments based on filter settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeReplicationTaskIndividualAssessmentsCmd).Standalone()

	dms_describeReplicationTaskIndividualAssessmentsCmd.Flags().String("filters", "", "Filters applied to the individual assessments described in the form of key-value pairs.")
	dms_describeReplicationTaskIndividualAssessmentsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	dms_describeReplicationTaskIndividualAssessmentsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dmsCmd.AddCommand(dms_describeReplicationTaskIndividualAssessmentsCmd)
}
