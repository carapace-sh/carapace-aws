package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getEvidenceFoldersByAssessmentCmd = &cobra.Command{
	Use:   "get-evidence-folders-by-assessment",
	Short: "Gets the evidence folders from a specified assessment in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getEvidenceFoldersByAssessmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_getEvidenceFoldersByAssessmentCmd).Standalone()

		auditmanager_getEvidenceFoldersByAssessmentCmd.Flags().String("assessment-id", "", "The unique identifier for the assessment.")
		auditmanager_getEvidenceFoldersByAssessmentCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
		auditmanager_getEvidenceFoldersByAssessmentCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		auditmanager_getEvidenceFoldersByAssessmentCmd.MarkFlagRequired("assessment-id")
	})
	auditmanagerCmd.AddCommand(auditmanager_getEvidenceFoldersByAssessmentCmd)
}
