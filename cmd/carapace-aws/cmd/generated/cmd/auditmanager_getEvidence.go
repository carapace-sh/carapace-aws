package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getEvidenceCmd = &cobra.Command{
	Use:   "get-evidence",
	Short: "Gets information about a specified evidence item.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getEvidenceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_getEvidenceCmd).Standalone()

		auditmanager_getEvidenceCmd.Flags().String("assessment-id", "", "The unique identifier for the assessment.")
		auditmanager_getEvidenceCmd.Flags().String("control-set-id", "", "The unique identifier for the control set.")
		auditmanager_getEvidenceCmd.Flags().String("evidence-folder-id", "", "The unique identifier for the folder that the evidence is stored in.")
		auditmanager_getEvidenceCmd.Flags().String("evidence-id", "", "The unique identifier for the evidence.")
		auditmanager_getEvidenceCmd.MarkFlagRequired("assessment-id")
		auditmanager_getEvidenceCmd.MarkFlagRequired("control-set-id")
		auditmanager_getEvidenceCmd.MarkFlagRequired("evidence-folder-id")
		auditmanager_getEvidenceCmd.MarkFlagRequired("evidence-id")
	})
	auditmanagerCmd.AddCommand(auditmanager_getEvidenceCmd)
}
