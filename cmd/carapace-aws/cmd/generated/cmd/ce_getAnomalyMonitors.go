package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getAnomalyMonitorsCmd = &cobra.Command{
	Use:   "get-anomaly-monitors",
	Short: "Retrieves the cost anomaly monitor definitions for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getAnomalyMonitorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_getAnomalyMonitorsCmd).Standalone()

		ce_getAnomalyMonitorsCmd.Flags().String("max-results", "", "The number of entries that a paginated response contains.")
		ce_getAnomalyMonitorsCmd.Flags().String("monitor-arn-list", "", "A list of cost anomaly monitor ARNs.")
		ce_getAnomalyMonitorsCmd.Flags().String("next-page-token", "", "The token to retrieve the next set of results.")
	})
	ceCmd.AddCommand(ce_getAnomalyMonitorsCmd)
}
