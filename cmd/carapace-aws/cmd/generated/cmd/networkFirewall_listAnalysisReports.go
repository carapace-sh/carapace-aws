package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_listAnalysisReportsCmd = &cobra.Command{
	Use:   "list-analysis-reports",
	Short: "Returns a list of all traffic analysis reports generated within the last 30 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_listAnalysisReportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_listAnalysisReportsCmd).Standalone()

		networkFirewall_listAnalysisReportsCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_listAnalysisReportsCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
		networkFirewall_listAnalysisReportsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Network Firewall to return for this request.")
		networkFirewall_listAnalysisReportsCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Network Firewall returns a `NextToken` value in the response.")
	})
	networkFirewallCmd.AddCommand(networkFirewall_listAnalysisReportsCmd)
}
