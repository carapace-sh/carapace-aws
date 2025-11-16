package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_createClassificationJobCmd = &cobra.Command{
	Use:   "create-classification-job",
	Short: "Creates and defines the settings for a classification job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_createClassificationJobCmd).Standalone()

	macie2_createClassificationJobCmd.Flags().String("allow-list-ids", "", "An array of unique identifiers, one for each allow list for the job to use when it analyzes data.")
	macie2_createClassificationJobCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure the idempotency of the request.")
	macie2_createClassificationJobCmd.Flags().String("custom-data-identifier-ids", "", "An array of unique identifiers, one for each custom data identifier for the job to use when it analyzes data.")
	macie2_createClassificationJobCmd.Flags().String("description", "", "A custom description of the job.")
	macie2_createClassificationJobCmd.Flags().String("initial-run", "", "For a recurring job, specifies whether to analyze all existing, eligible objects immediately after the job is created (true).")
	macie2_createClassificationJobCmd.Flags().String("job-type", "", "The schedule for running the job.")
	macie2_createClassificationJobCmd.Flags().String("managed-data-identifier-ids", "", "An array of unique identifiers, one for each managed data identifier for the job to include (use) or exclude (not use) when it analyzes data.")
	macie2_createClassificationJobCmd.Flags().String("managed-data-identifier-selector", "", "The selection type to apply when determining which managed data identifiers the job uses to analyze data.")
	macie2_createClassificationJobCmd.Flags().String("name", "", "A custom name for the job.")
	macie2_createClassificationJobCmd.Flags().String("s3-job-definition", "", "The S3 buckets that contain the objects to analyze, and the scope of that analysis.")
	macie2_createClassificationJobCmd.Flags().String("sampling-percentage", "", "The sampling depth, as a percentage, for the job to apply when processing objects.")
	macie2_createClassificationJobCmd.Flags().String("schedule-frequency", "", "The recurrence pattern for running the job.")
	macie2_createClassificationJobCmd.Flags().String("tags", "", "A map of key-value pairs that specifies the tags to associate with the job.")
	macie2_createClassificationJobCmd.MarkFlagRequired("client-token")
	macie2_createClassificationJobCmd.MarkFlagRequired("job-type")
	macie2_createClassificationJobCmd.MarkFlagRequired("name")
	macie2_createClassificationJobCmd.MarkFlagRequired("s3-job-definition")
	macie2Cmd.AddCommand(macie2_createClassificationJobCmd)
}
