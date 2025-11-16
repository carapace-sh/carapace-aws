package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_startAnalysisReportCmd = &cobra.Command{
	Use:   "start-analysis-report",
	Short: "Generates a traffic analysis report for the timeframe and traffic type you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_startAnalysisReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_startAnalysisReportCmd).Standalone()

		networkFirewall_startAnalysisReportCmd.Flags().String("analysis-type", "", "The type of traffic that will be used to generate a report.")
		networkFirewall_startAnalysisReportCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_startAnalysisReportCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
		networkFirewall_startAnalysisReportCmd.MarkFlagRequired("analysis-type")
	})
	networkFirewallCmd.AddCommand(networkFirewall_startAnalysisReportCmd)
}
