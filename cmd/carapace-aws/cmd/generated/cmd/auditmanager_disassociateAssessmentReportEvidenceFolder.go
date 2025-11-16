package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_disassociateAssessmentReportEvidenceFolderCmd = &cobra.Command{
	Use:   "disassociate-assessment-report-evidence-folder",
	Short: "Disassociates an evidence folder from the specified assessment report in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_disassociateAssessmentReportEvidenceFolderCmd).Standalone()

	auditmanager_disassociateAssessmentReportEvidenceFolderCmd.Flags().String("assessment-id", "", "The unique identifier for the assessment.")
	auditmanager_disassociateAssessmentReportEvidenceFolderCmd.Flags().String("evidence-folder-id", "", "The unique identifier for the folder that the evidence is stored in.")
	auditmanager_disassociateAssessmentReportEvidenceFolderCmd.MarkFlagRequired("assessment-id")
	auditmanager_disassociateAssessmentReportEvidenceFolderCmd.MarkFlagRequired("evidence-folder-id")
	auditmanagerCmd.AddCommand(auditmanager_disassociateAssessmentReportEvidenceFolderCmd)
}
