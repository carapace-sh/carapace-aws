package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_batchDisassociateAssessmentReportEvidenceCmd = &cobra.Command{
	Use:   "batch-disassociate-assessment-report-evidence",
	Short: "Disassociates a list of evidence from an assessment report in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_batchDisassociateAssessmentReportEvidenceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_batchDisassociateAssessmentReportEvidenceCmd).Standalone()

		auditmanager_batchDisassociateAssessmentReportEvidenceCmd.Flags().String("assessment-id", "", "The identifier for the assessment.")
		auditmanager_batchDisassociateAssessmentReportEvidenceCmd.Flags().String("evidence-folder-id", "", "The identifier for the folder that the evidence is stored in.")
		auditmanager_batchDisassociateAssessmentReportEvidenceCmd.Flags().String("evidence-ids", "", "The list of evidence identifiers.")
		auditmanager_batchDisassociateAssessmentReportEvidenceCmd.MarkFlagRequired("assessment-id")
		auditmanager_batchDisassociateAssessmentReportEvidenceCmd.MarkFlagRequired("evidence-folder-id")
		auditmanager_batchDisassociateAssessmentReportEvidenceCmd.MarkFlagRequired("evidence-ids")
	})
	auditmanagerCmd.AddCommand(auditmanager_batchDisassociateAssessmentReportEvidenceCmd)
}
