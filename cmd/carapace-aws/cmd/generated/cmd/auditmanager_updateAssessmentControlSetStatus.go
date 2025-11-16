package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_updateAssessmentControlSetStatusCmd = &cobra.Command{
	Use:   "update-assessment-control-set-status",
	Short: "Updates the status of a control set in an Audit Manager assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_updateAssessmentControlSetStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_updateAssessmentControlSetStatusCmd).Standalone()

		auditmanager_updateAssessmentControlSetStatusCmd.Flags().String("assessment-id", "", "The unique identifier for the assessment.")
		auditmanager_updateAssessmentControlSetStatusCmd.Flags().String("comment", "", "The comment that's related to the status update.")
		auditmanager_updateAssessmentControlSetStatusCmd.Flags().String("control-set-id", "", "The unique identifier for the control set.")
		auditmanager_updateAssessmentControlSetStatusCmd.Flags().String("status", "", "The status of the control set that's being updated.")
		auditmanager_updateAssessmentControlSetStatusCmd.MarkFlagRequired("assessment-id")
		auditmanager_updateAssessmentControlSetStatusCmd.MarkFlagRequired("comment")
		auditmanager_updateAssessmentControlSetStatusCmd.MarkFlagRequired("control-set-id")
		auditmanager_updateAssessmentControlSetStatusCmd.MarkFlagRequired("status")
	})
	auditmanagerCmd.AddCommand(auditmanager_updateAssessmentControlSetStatusCmd)
}
