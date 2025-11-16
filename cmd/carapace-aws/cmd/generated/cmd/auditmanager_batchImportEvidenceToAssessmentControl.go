package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_batchImportEvidenceToAssessmentControlCmd = &cobra.Command{
	Use:   "batch-import-evidence-to-assessment-control",
	Short: "Adds one or more pieces of evidence to a control in an Audit Manager assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_batchImportEvidenceToAssessmentControlCmd).Standalone()

	auditmanager_batchImportEvidenceToAssessmentControlCmd.Flags().String("assessment-id", "", "The identifier for the assessment.")
	auditmanager_batchImportEvidenceToAssessmentControlCmd.Flags().String("control-id", "", "The identifier for the control.")
	auditmanager_batchImportEvidenceToAssessmentControlCmd.Flags().String("control-set-id", "", "The identifier for the control set.")
	auditmanager_batchImportEvidenceToAssessmentControlCmd.Flags().String("manual-evidence", "", "The list of manual evidence objects.")
	auditmanager_batchImportEvidenceToAssessmentControlCmd.MarkFlagRequired("assessment-id")
	auditmanager_batchImportEvidenceToAssessmentControlCmd.MarkFlagRequired("control-id")
	auditmanager_batchImportEvidenceToAssessmentControlCmd.MarkFlagRequired("control-set-id")
	auditmanager_batchImportEvidenceToAssessmentControlCmd.MarkFlagRequired("manual-evidence")
	auditmanagerCmd.AddCommand(auditmanager_batchImportEvidenceToAssessmentControlCmd)
}
