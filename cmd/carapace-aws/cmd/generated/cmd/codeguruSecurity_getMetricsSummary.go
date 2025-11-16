package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruSecurity_getMetricsSummaryCmd = &cobra.Command{
	Use:   "get-metrics-summary",
	Short: "Returns a summary of metrics for an account from a specified date, including number of open findings, the categories with most findings, the scans with most open findings, and scans with most open critical findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruSecurity_getMetricsSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruSecurity_getMetricsSummaryCmd).Standalone()

		codeguruSecurity_getMetricsSummaryCmd.Flags().String("date", "", "The date you want to retrieve summary metrics from, rounded to the nearest day.")
		codeguruSecurity_getMetricsSummaryCmd.MarkFlagRequired("date")
	})
	codeguruSecurityCmd.AddCommand(codeguruSecurity_getMetricsSummaryCmd)
}
