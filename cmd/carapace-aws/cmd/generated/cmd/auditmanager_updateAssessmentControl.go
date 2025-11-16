package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_updateAssessmentControlCmd = &cobra.Command{
	Use:   "update-assessment-control",
	Short: "Updates a control within an assessment in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_updateAssessmentControlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_updateAssessmentControlCmd).Standalone()

		auditmanager_updateAssessmentControlCmd.Flags().String("assessment-id", "", "The unique identifier for the assessment.")
		auditmanager_updateAssessmentControlCmd.Flags().String("comment-body", "", "The comment body text for the control.")
		auditmanager_updateAssessmentControlCmd.Flags().String("control-id", "", "The unique identifier for the control.")
		auditmanager_updateAssessmentControlCmd.Flags().String("control-set-id", "", "The unique identifier for the control set.")
		auditmanager_updateAssessmentControlCmd.Flags().String("control-status", "", "The status of the control.")
		auditmanager_updateAssessmentControlCmd.MarkFlagRequired("assessment-id")
		auditmanager_updateAssessmentControlCmd.MarkFlagRequired("control-id")
		auditmanager_updateAssessmentControlCmd.MarkFlagRequired("control-set-id")
	})
	auditmanagerCmd.AddCommand(auditmanager_updateAssessmentControlCmd)
}
