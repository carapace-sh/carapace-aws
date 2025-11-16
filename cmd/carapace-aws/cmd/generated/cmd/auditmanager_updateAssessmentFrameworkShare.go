package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_updateAssessmentFrameworkShareCmd = &cobra.Command{
	Use:   "update-assessment-framework-share",
	Short: "Updates a share request for a custom framework in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_updateAssessmentFrameworkShareCmd).Standalone()

	auditmanager_updateAssessmentFrameworkShareCmd.Flags().String("action", "", "Specifies the update action for the share request.")
	auditmanager_updateAssessmentFrameworkShareCmd.Flags().String("request-id", "", "The unique identifier for the share request.")
	auditmanager_updateAssessmentFrameworkShareCmd.Flags().String("request-type", "", "Specifies whether the share request is a sent request or a received request.")
	auditmanager_updateAssessmentFrameworkShareCmd.MarkFlagRequired("action")
	auditmanager_updateAssessmentFrameworkShareCmd.MarkFlagRequired("request-id")
	auditmanager_updateAssessmentFrameworkShareCmd.MarkFlagRequired("request-type")
	auditmanagerCmd.AddCommand(auditmanager_updateAssessmentFrameworkShareCmd)
}
