package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_updateAssessmentCmd = &cobra.Command{
	Use:   "update-assessment",
	Short: "Edits an Audit Manager assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_updateAssessmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_updateAssessmentCmd).Standalone()

		auditmanager_updateAssessmentCmd.Flags().String("assessment-description", "", "The description of the assessment.")
		auditmanager_updateAssessmentCmd.Flags().String("assessment-id", "", "The unique identifier for the assessment.")
		auditmanager_updateAssessmentCmd.Flags().String("assessment-name", "", "The name of the assessment to be updated.")
		auditmanager_updateAssessmentCmd.Flags().String("assessment-reports-destination", "", "The assessment report storage destination for the assessment that's being updated.")
		auditmanager_updateAssessmentCmd.Flags().String("roles", "", "The list of roles for the assessment.")
		auditmanager_updateAssessmentCmd.Flags().String("scope", "", "The scope of the assessment.")
		auditmanager_updateAssessmentCmd.MarkFlagRequired("assessment-id")
		auditmanager_updateAssessmentCmd.MarkFlagRequired("scope")
	})
	auditmanagerCmd.AddCommand(auditmanager_updateAssessmentCmd)
}
