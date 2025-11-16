package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_deleteAssessmentReportCmd = &cobra.Command{
	Use:   "delete-assessment-report",
	Short: "Deletes an assessment report in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_deleteAssessmentReportCmd).Standalone()

	auditmanager_deleteAssessmentReportCmd.Flags().String("assessment-id", "", "The unique identifier for the assessment.")
	auditmanager_deleteAssessmentReportCmd.Flags().String("assessment-report-id", "", "The unique identifier for the assessment report.")
	auditmanager_deleteAssessmentReportCmd.MarkFlagRequired("assessment-id")
	auditmanager_deleteAssessmentReportCmd.MarkFlagRequired("assessment-report-id")
	auditmanagerCmd.AddCommand(auditmanager_deleteAssessmentReportCmd)
}
