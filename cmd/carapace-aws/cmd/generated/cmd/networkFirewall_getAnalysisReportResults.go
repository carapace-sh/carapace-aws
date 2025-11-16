package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_getAnalysisReportResultsCmd = &cobra.Command{
	Use:   "get-analysis-report-results",
	Short: "The results of a `COMPLETED` analysis report generated with [StartAnalysisReport]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_getAnalysisReportResultsCmd).Standalone()

	networkFirewall_getAnalysisReportResultsCmd.Flags().String("analysis-report-id", "", "The unique ID of the query that ran when you requested an analysis report.")
	networkFirewall_getAnalysisReportResultsCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
	networkFirewall_getAnalysisReportResultsCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
	networkFirewall_getAnalysisReportResultsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Network Firewall to return for this request.")
	networkFirewall_getAnalysisReportResultsCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Network Firewall returns a `NextToken` value in the response.")
	networkFirewall_getAnalysisReportResultsCmd.MarkFlagRequired("analysis-report-id")
	networkFirewallCmd.AddCommand(networkFirewall_getAnalysisReportResultsCmd)
}
