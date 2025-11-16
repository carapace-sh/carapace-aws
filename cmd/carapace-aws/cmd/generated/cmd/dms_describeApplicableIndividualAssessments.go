package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeApplicableIndividualAssessmentsCmd = &cobra.Command{
	Use:   "describe-applicable-individual-assessments",
	Short: "Provides a list of individual assessments that you can specify for a new premigration assessment run, given one or more parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeApplicableIndividualAssessmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeApplicableIndividualAssessmentsCmd).Standalone()

		dms_describeApplicableIndividualAssessmentsCmd.Flags().String("marker", "", "Optional pagination token provided by a previous request.")
		dms_describeApplicableIndividualAssessmentsCmd.Flags().String("max-records", "", "Maximum number of records to include in the response.")
		dms_describeApplicableIndividualAssessmentsCmd.Flags().String("migration-type", "", "Name of the migration type that each provided individual assessment must support.")
		dms_describeApplicableIndividualAssessmentsCmd.Flags().String("replication-config-arn", "", "Amazon Resource Name (ARN) of a serverless replication on which you want to base the default list of individual assessments.")
		dms_describeApplicableIndividualAssessmentsCmd.Flags().String("replication-instance-arn", "", "ARN of a replication instance on which you want to base the default list of individual assessments.")
		dms_describeApplicableIndividualAssessmentsCmd.Flags().String("replication-task-arn", "", "Amazon Resource Name (ARN) of a migration task on which you want to base the default list of individual assessments.")
		dms_describeApplicableIndividualAssessmentsCmd.Flags().String("source-engine-name", "", "Name of a database engine that the specified replication instance supports as a source.")
		dms_describeApplicableIndividualAssessmentsCmd.Flags().String("target-engine-name", "", "Name of a database engine that the specified replication instance supports as a target.")
	})
	dmsCmd.AddCommand(dms_describeApplicableIndividualAssessmentsCmd)
}
