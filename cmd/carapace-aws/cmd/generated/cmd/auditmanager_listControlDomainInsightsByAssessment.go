package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_listControlDomainInsightsByAssessmentCmd = &cobra.Command{
	Use:   "list-control-domain-insights-by-assessment",
	Short: "Lists analytics data for control domains within a specified active assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_listControlDomainInsightsByAssessmentCmd).Standalone()

	auditmanager_listControlDomainInsightsByAssessmentCmd.Flags().String("assessment-id", "", "The unique identifier for the active assessment.")
	auditmanager_listControlDomainInsightsByAssessmentCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
	auditmanager_listControlDomainInsightsByAssessmentCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	auditmanager_listControlDomainInsightsByAssessmentCmd.MarkFlagRequired("assessment-id")
	auditmanagerCmd.AddCommand(auditmanager_listControlDomainInsightsByAssessmentCmd)
}
