package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listWhatIfAnalysesCmd = &cobra.Command{
	Use:   "list-what-if-analyses",
	Short: "Returns a list of what-if analyses created using the [CreateWhatIfAnalysis]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listWhatIfAnalysesCmd).Standalone()

	forecast_listWhatIfAnalysesCmd.Flags().String("filters", "", "An array of filters.")
	forecast_listWhatIfAnalysesCmd.Flags().String("max-results", "", "The number of items to return in the response.")
	forecast_listWhatIfAnalysesCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a `NextToken`.")
	forecastCmd.AddCommand(forecast_listWhatIfAnalysesCmd)
}
