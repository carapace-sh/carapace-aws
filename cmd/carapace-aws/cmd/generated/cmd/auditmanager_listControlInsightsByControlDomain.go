package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_listControlInsightsByControlDomainCmd = &cobra.Command{
	Use:   "list-control-insights-by-control-domain",
	Short: "Lists the latest analytics data for controls within a specific control domain across all active assessments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_listControlInsightsByControlDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_listControlInsightsByControlDomainCmd).Standalone()

		auditmanager_listControlInsightsByControlDomainCmd.Flags().String("control-domain-id", "", "The unique identifier for the control domain.")
		auditmanager_listControlInsightsByControlDomainCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
		auditmanager_listControlInsightsByControlDomainCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		auditmanager_listControlInsightsByControlDomainCmd.MarkFlagRequired("control-domain-id")
	})
	auditmanagerCmd.AddCommand(auditmanager_listControlInsightsByControlDomainCmd)
}
