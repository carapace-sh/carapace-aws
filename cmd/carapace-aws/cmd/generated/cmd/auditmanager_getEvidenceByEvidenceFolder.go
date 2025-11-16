package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getEvidenceByEvidenceFolderCmd = &cobra.Command{
	Use:   "get-evidence-by-evidence-folder",
	Short: "Gets all evidence from a specified evidence folder in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getEvidenceByEvidenceFolderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_getEvidenceByEvidenceFolderCmd).Standalone()

		auditmanager_getEvidenceByEvidenceFolderCmd.Flags().String("assessment-id", "", "The identifier for the assessment.")
		auditmanager_getEvidenceByEvidenceFolderCmd.Flags().String("control-set-id", "", "The identifier for the control set.")
		auditmanager_getEvidenceByEvidenceFolderCmd.Flags().String("evidence-folder-id", "", "The unique identifier for the folder that the evidence is stored in.")
		auditmanager_getEvidenceByEvidenceFolderCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
		auditmanager_getEvidenceByEvidenceFolderCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		auditmanager_getEvidenceByEvidenceFolderCmd.MarkFlagRequired("assessment-id")
		auditmanager_getEvidenceByEvidenceFolderCmd.MarkFlagRequired("control-set-id")
		auditmanager_getEvidenceByEvidenceFolderCmd.MarkFlagRequired("evidence-folder-id")
	})
	auditmanagerCmd.AddCommand(auditmanager_getEvidenceByEvidenceFolderCmd)
}
