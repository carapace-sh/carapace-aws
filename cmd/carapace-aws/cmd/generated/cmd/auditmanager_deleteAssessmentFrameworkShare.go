package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_deleteAssessmentFrameworkShareCmd = &cobra.Command{
	Use:   "delete-assessment-framework-share",
	Short: "Deletes a share request for a custom framework in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_deleteAssessmentFrameworkShareCmd).Standalone()

	auditmanager_deleteAssessmentFrameworkShareCmd.Flags().String("request-id", "", "The unique identifier for the share request to be deleted.")
	auditmanager_deleteAssessmentFrameworkShareCmd.Flags().String("request-type", "", "Specifies whether the share request is a sent request or a received request.")
	auditmanager_deleteAssessmentFrameworkShareCmd.MarkFlagRequired("request-id")
	auditmanager_deleteAssessmentFrameworkShareCmd.MarkFlagRequired("request-type")
	auditmanagerCmd.AddCommand(auditmanager_deleteAssessmentFrameworkShareCmd)
}
