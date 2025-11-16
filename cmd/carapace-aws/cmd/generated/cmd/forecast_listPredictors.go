package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listPredictorsCmd = &cobra.Command{
	Use:   "list-predictors",
	Short: "Returns a list of predictors created using the [CreateAutoPredictor]() or [CreatePredictor]() operations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listPredictorsCmd).Standalone()

	forecast_listPredictorsCmd.Flags().String("filters", "", "An array of filters.")
	forecast_listPredictorsCmd.Flags().String("max-results", "", "The number of items to return in the response.")
	forecast_listPredictorsCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a `NextToken`.")
	forecastCmd.AddCommand(forecast_listPredictorsCmd)
}
