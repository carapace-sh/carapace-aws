package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_validateAssessmentReportIntegrityCmd = &cobra.Command{
	Use:   "validate-assessment-report-integrity",
	Short: "Validates the integrity of an assessment report in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_validateAssessmentReportIntegrityCmd).Standalone()

	auditmanager_validateAssessmentReportIntegrityCmd.Flags().String("s3-relative-path", "", "The relative path of the Amazon S3 bucket that the assessment report is stored in.")
	auditmanager_validateAssessmentReportIntegrityCmd.MarkFlagRequired("s3-relative-path")
	auditmanagerCmd.AddCommand(auditmanager_validateAssessmentReportIntegrityCmd)
}
