package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_updateAssessmentStatusCmd = &cobra.Command{
	Use:   "update-assessment-status",
	Short: "Updates the status of an assessment in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_updateAssessmentStatusCmd).Standalone()

	auditmanager_updateAssessmentStatusCmd.Flags().String("assessment-id", "", "The unique identifier for the assessment.")
	auditmanager_updateAssessmentStatusCmd.Flags().String("status", "", "The current status of the assessment.")
	auditmanager_updateAssessmentStatusCmd.MarkFlagRequired("assessment-id")
	auditmanager_updateAssessmentStatusCmd.MarkFlagRequired("status")
	auditmanagerCmd.AddCommand(auditmanager_updateAssessmentStatusCmd)
}
