package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_createAssessmentReportCmd = &cobra.Command{
	Use:   "create-assessment-report",
	Short: "Creates an assessment report for the specified assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_createAssessmentReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_createAssessmentReportCmd).Standalone()

		auditmanager_createAssessmentReportCmd.Flags().String("assessment-id", "", "The identifier for the assessment.")
		auditmanager_createAssessmentReportCmd.Flags().String("description", "", "The description of the assessment report.")
		auditmanager_createAssessmentReportCmd.Flags().String("name", "", "The name of the new assessment report.")
		auditmanager_createAssessmentReportCmd.Flags().String("query-statement", "", "A SQL statement that represents an evidence finder query.")
		auditmanager_createAssessmentReportCmd.MarkFlagRequired("assessment-id")
		auditmanager_createAssessmentReportCmd.MarkFlagRequired("name")
	})
	auditmanagerCmd.AddCommand(auditmanager_createAssessmentReportCmd)
}
