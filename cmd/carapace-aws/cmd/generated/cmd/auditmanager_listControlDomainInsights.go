package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_listControlDomainInsightsCmd = &cobra.Command{
	Use:   "list-control-domain-insights",
	Short: "Lists the latest analytics data for control domains across all of your active assessments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_listControlDomainInsightsCmd).Standalone()

	auditmanager_listControlDomainInsightsCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
	auditmanager_listControlDomainInsightsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	auditmanagerCmd.AddCommand(auditmanager_listControlDomainInsightsCmd)
}
