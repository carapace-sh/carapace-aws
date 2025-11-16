package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getEvidenceFoldersByAssessmentControlCmd = &cobra.Command{
	Use:   "get-evidence-folders-by-assessment-control",
	Short: "Gets a list of evidence folders that are associated with a specified control in an Audit Manager assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getEvidenceFoldersByAssessmentControlCmd).Standalone()

	auditmanager_getEvidenceFoldersByAssessmentControlCmd.Flags().String("assessment-id", "", "The identifier for the assessment.")
	auditmanager_getEvidenceFoldersByAssessmentControlCmd.Flags().String("control-id", "", "The identifier for the control.")
	auditmanager_getEvidenceFoldersByAssessmentControlCmd.Flags().String("control-set-id", "", "The identifier for the control set.")
	auditmanager_getEvidenceFoldersByAssessmentControlCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
	auditmanager_getEvidenceFoldersByAssessmentControlCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	auditmanager_getEvidenceFoldersByAssessmentControlCmd.MarkFlagRequired("assessment-id")
	auditmanager_getEvidenceFoldersByAssessmentControlCmd.MarkFlagRequired("control-id")
	auditmanager_getEvidenceFoldersByAssessmentControlCmd.MarkFlagRequired("control-set-id")
	auditmanagerCmd.AddCommand(auditmanager_getEvidenceFoldersByAssessmentControlCmd)
}
