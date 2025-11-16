package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getInsightsByAssessmentCmd = &cobra.Command{
	Use:   "get-insights-by-assessment",
	Short: "Gets the latest analytics data for a specific active assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getInsightsByAssessmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_getInsightsByAssessmentCmd).Standalone()

		auditmanager_getInsightsByAssessmentCmd.Flags().String("assessment-id", "", "The unique identifier for the assessment.")
		auditmanager_getInsightsByAssessmentCmd.MarkFlagRequired("assessment-id")
	})
	auditmanagerCmd.AddCommand(auditmanager_getInsightsByAssessmentCmd)
}
