package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_batchAssociateAssessmentReportEvidenceCmd = &cobra.Command{
	Use:   "batch-associate-assessment-report-evidence",
	Short: "Associates a list of evidence to an assessment report in an Audit Manager assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_batchAssociateAssessmentReportEvidenceCmd).Standalone()

	auditmanager_batchAssociateAssessmentReportEvidenceCmd.Flags().String("assessment-id", "", "The identifier for the assessment.")
	auditmanager_batchAssociateAssessmentReportEvidenceCmd.Flags().String("evidence-folder-id", "", "The identifier for the folder that the evidence is stored in.")
	auditmanager_batchAssociateAssessmentReportEvidenceCmd.Flags().String("evidence-ids", "", "The list of evidence identifiers.")
	auditmanager_batchAssociateAssessmentReportEvidenceCmd.MarkFlagRequired("assessment-id")
	auditmanager_batchAssociateAssessmentReportEvidenceCmd.MarkFlagRequired("evidence-folder-id")
	auditmanager_batchAssociateAssessmentReportEvidenceCmd.MarkFlagRequired("evidence-ids")
	auditmanagerCmd.AddCommand(auditmanager_batchAssociateAssessmentReportEvidenceCmd)
}
