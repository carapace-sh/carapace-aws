package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getEvidenceFolderCmd = &cobra.Command{
	Use:   "get-evidence-folder",
	Short: "Gets an evidence folder from a specified assessment in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getEvidenceFolderCmd).Standalone()

	auditmanager_getEvidenceFolderCmd.Flags().String("assessment-id", "", "The unique identifier for the assessment.")
	auditmanager_getEvidenceFolderCmd.Flags().String("control-set-id", "", "The unique identifier for the control set.")
	auditmanager_getEvidenceFolderCmd.Flags().String("evidence-folder-id", "", "The unique identifier for the folder that the evidence is stored in.")
	auditmanager_getEvidenceFolderCmd.MarkFlagRequired("assessment-id")
	auditmanager_getEvidenceFolderCmd.MarkFlagRequired("control-set-id")
	auditmanager_getEvidenceFolderCmd.MarkFlagRequired("evidence-folder-id")
	auditmanagerCmd.AddCommand(auditmanager_getEvidenceFolderCmd)
}
