package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_startReplicationTaskAssessmentRunCmd = &cobra.Command{
	Use:   "start-replication-task-assessment-run",
	Short: "Starts a new premigration assessment run for one or more individual assessments of a migration task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_startReplicationTaskAssessmentRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_startReplicationTaskAssessmentRunCmd).Standalone()

		dms_startReplicationTaskAssessmentRunCmd.Flags().String("assessment-run-name", "", "Unique name to identify the assessment run.")
		dms_startReplicationTaskAssessmentRunCmd.Flags().String("exclude", "", "Space-separated list of names for specific individual assessments that you want to exclude.")
		dms_startReplicationTaskAssessmentRunCmd.Flags().String("include-only", "", "Space-separated list of names for specific individual assessments that you want to include.")
		dms_startReplicationTaskAssessmentRunCmd.Flags().String("replication-task-arn", "", "Amazon Resource Name (ARN) of the migration task associated with the premigration assessment run that you want to start.")
		dms_startReplicationTaskAssessmentRunCmd.Flags().String("result-encryption-mode", "", "Encryption mode that you can specify to encrypt the results of this assessment run.")
		dms_startReplicationTaskAssessmentRunCmd.Flags().String("result-kms-key-arn", "", "ARN of a custom KMS encryption key that you specify when you set `ResultEncryptionMode` to `\"SSE_KMS`\".")
		dms_startReplicationTaskAssessmentRunCmd.Flags().String("result-location-bucket", "", "Amazon S3 bucket where you want DMS to store the results of this assessment run.")
		dms_startReplicationTaskAssessmentRunCmd.Flags().String("result-location-folder", "", "Folder within an Amazon S3 bucket where you want DMS to store the results of this assessment run.")
		dms_startReplicationTaskAssessmentRunCmd.Flags().String("service-access-role-arn", "", "ARN of the service role needed to start the assessment run.")
		dms_startReplicationTaskAssessmentRunCmd.Flags().String("tags", "", "One or more tags to be assigned to the premigration assessment run that you want to start.")
		dms_startReplicationTaskAssessmentRunCmd.MarkFlagRequired("assessment-run-name")
		dms_startReplicationTaskAssessmentRunCmd.MarkFlagRequired("replication-task-arn")
		dms_startReplicationTaskAssessmentRunCmd.MarkFlagRequired("result-location-bucket")
		dms_startReplicationTaskAssessmentRunCmd.MarkFlagRequired("service-access-role-arn")
	})
	dmsCmd.AddCommand(dms_startReplicationTaskAssessmentRunCmd)
}
