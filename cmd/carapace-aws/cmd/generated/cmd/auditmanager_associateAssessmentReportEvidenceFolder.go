package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_associateAssessmentReportEvidenceFolderCmd = &cobra.Command{
	Use:   "associate-assessment-report-evidence-folder",
	Short: "Associates an evidence folder to an assessment report in an Audit Manager assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_associateAssessmentReportEvidenceFolderCmd).Standalone()

	auditmanager_associateAssessmentReportEvidenceFolderCmd.Flags().String("assessment-id", "", "The identifier for the assessment.")
	auditmanager_associateAssessmentReportEvidenceFolderCmd.Flags().String("evidence-folder-id", "", "The identifier for the folder that the evidence is stored in.")
	auditmanager_associateAssessmentReportEvidenceFolderCmd.MarkFlagRequired("assessment-id")
	auditmanager_associateAssessmentReportEvidenceFolderCmd.MarkFlagRequired("evidence-folder-id")
	auditmanagerCmd.AddCommand(auditmanager_associateAssessmentReportEvidenceFolderCmd)
}
