package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getAssessmentCmd = &cobra.Command{
	Use:   "get-assessment",
	Short: "Gets information about a specified assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getAssessmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_getAssessmentCmd).Standalone()

		auditmanager_getAssessmentCmd.Flags().String("assessment-id", "", "The unique identifier for the assessment.")
		auditmanager_getAssessmentCmd.MarkFlagRequired("assessment-id")
	})
	auditmanagerCmd.AddCommand(auditmanager_getAssessmentCmd)
}
