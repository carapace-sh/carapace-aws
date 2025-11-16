package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listExplainabilityExportsCmd = &cobra.Command{
	Use:   "list-explainability-exports",
	Short: "Returns a list of Explainability exports created using the [CreateExplainabilityExport]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listExplainabilityExportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_listExplainabilityExportsCmd).Standalone()

		forecast_listExplainabilityExportsCmd.Flags().String("filters", "", "An array of filters.")
		forecast_listExplainabilityExportsCmd.Flags().String("max-results", "", "The number of items to return in the response.")
		forecast_listExplainabilityExportsCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a NextToken.")
	})
	forecastCmd.AddCommand(forecast_listExplainabilityExportsCmd)
}
