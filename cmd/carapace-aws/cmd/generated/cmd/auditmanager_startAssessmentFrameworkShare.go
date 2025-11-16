package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_startAssessmentFrameworkShareCmd = &cobra.Command{
	Use:   "start-assessment-framework-share",
	Short: "Creates a share request for a custom framework in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_startAssessmentFrameworkShareCmd).Standalone()

	auditmanager_startAssessmentFrameworkShareCmd.Flags().String("comment", "", "An optional comment from the sender about the share request.")
	auditmanager_startAssessmentFrameworkShareCmd.Flags().String("destination-account", "", "The Amazon Web Services account of the recipient.")
	auditmanager_startAssessmentFrameworkShareCmd.Flags().String("destination-region", "", "The Amazon Web Services Region of the recipient.")
	auditmanager_startAssessmentFrameworkShareCmd.Flags().String("framework-id", "", "The unique identifier for the custom framework to be shared.")
	auditmanager_startAssessmentFrameworkShareCmd.MarkFlagRequired("destination-account")
	auditmanager_startAssessmentFrameworkShareCmd.MarkFlagRequired("destination-region")
	auditmanager_startAssessmentFrameworkShareCmd.MarkFlagRequired("framework-id")
	auditmanagerCmd.AddCommand(auditmanager_startAssessmentFrameworkShareCmd)
}
