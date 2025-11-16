package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getAssessmentReportUrlCmd = &cobra.Command{
	Use:   "get-assessment-report-url",
	Short: "Gets the URL of an assessment report in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getAssessmentReportUrlCmd).Standalone()

	auditmanager_getAssessmentReportUrlCmd.Flags().String("assessment-id", "", "The unique identifier for the assessment.")
	auditmanager_getAssessmentReportUrlCmd.Flags().String("assessment-report-id", "", "The unique identifier for the assessment report.")
	auditmanager_getAssessmentReportUrlCmd.MarkFlagRequired("assessment-id")
	auditmanager_getAssessmentReportUrlCmd.MarkFlagRequired("assessment-report-id")
	auditmanagerCmd.AddCommand(auditmanager_getAssessmentReportUrlCmd)
}
