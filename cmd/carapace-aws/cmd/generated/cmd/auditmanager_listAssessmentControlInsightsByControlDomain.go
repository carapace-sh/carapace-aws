package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_listAssessmentControlInsightsByControlDomainCmd = &cobra.Command{
	Use:   "list-assessment-control-insights-by-control-domain",
	Short: "Lists the latest analytics data for controls within a specific control domain and a specific active assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_listAssessmentControlInsightsByControlDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_listAssessmentControlInsightsByControlDomainCmd).Standalone()

		auditmanager_listAssessmentControlInsightsByControlDomainCmd.Flags().String("assessment-id", "", "The unique identifier for the active assessment.")
		auditmanager_listAssessmentControlInsightsByControlDomainCmd.Flags().String("control-domain-id", "", "The unique identifier for the control domain.")
		auditmanager_listAssessmentControlInsightsByControlDomainCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
		auditmanager_listAssessmentControlInsightsByControlDomainCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		auditmanager_listAssessmentControlInsightsByControlDomainCmd.MarkFlagRequired("assessment-id")
		auditmanager_listAssessmentControlInsightsByControlDomainCmd.MarkFlagRequired("control-domain-id")
	})
	auditmanagerCmd.AddCommand(auditmanager_listAssessmentControlInsightsByControlDomainCmd)
}
