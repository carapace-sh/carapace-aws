package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_deleteAssessmentCmd = &cobra.Command{
	Use:   "delete-assessment",
	Short: "Deletes an assessment in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_deleteAssessmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_deleteAssessmentCmd).Standalone()

		auditmanager_deleteAssessmentCmd.Flags().String("assessment-id", "", "The identifier for the assessment.")
		auditmanager_deleteAssessmentCmd.MarkFlagRequired("assessment-id")
	})
	auditmanagerCmd.AddCommand(auditmanager_deleteAssessmentCmd)
}
