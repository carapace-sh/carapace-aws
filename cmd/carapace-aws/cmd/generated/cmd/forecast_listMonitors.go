package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listMonitorsCmd = &cobra.Command{
	Use:   "list-monitors",
	Short: "Returns a list of monitors created with the [CreateMonitor]() operation and [CreateAutoPredictor]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listMonitorsCmd).Standalone()

	forecast_listMonitorsCmd.Flags().String("filters", "", "An array of filters.")
	forecast_listMonitorsCmd.Flags().String("max-results", "", "The maximum number of monitors to include in the response.")
	forecast_listMonitorsCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a `NextToken`.")
	forecastCmd.AddCommand(forecast_listMonitorsCmd)
}
