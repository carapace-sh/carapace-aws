package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruSecurity_listFindingsMetricsCmd = &cobra.Command{
	Use:   "list-findings-metrics",
	Short: "Returns metrics about all findings in an account within a specified time range.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruSecurity_listFindingsMetricsCmd).Standalone()

	codeguruSecurity_listFindingsMetricsCmd.Flags().String("end-date", "", "The end date of the interval which you want to retrieve metrics from.")
	codeguruSecurity_listFindingsMetricsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	codeguruSecurity_listFindingsMetricsCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	codeguruSecurity_listFindingsMetricsCmd.Flags().String("start-date", "", "The start date of the interval which you want to retrieve metrics from.")
	codeguruSecurity_listFindingsMetricsCmd.MarkFlagRequired("end-date")
	codeguruSecurity_listFindingsMetricsCmd.MarkFlagRequired("start-date")
	codeguruSecurityCmd.AddCommand(codeguruSecurity_listFindingsMetricsCmd)
}
