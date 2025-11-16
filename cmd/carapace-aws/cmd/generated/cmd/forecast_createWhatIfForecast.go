package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_createWhatIfForecastCmd = &cobra.Command{
	Use:   "create-what-if-forecast",
	Short: "A what-if forecast is a forecast that is created from a modified version of the baseline forecast.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_createWhatIfForecastCmd).Standalone()

	forecast_createWhatIfForecastCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/forecast/latest/dg/tagging-forecast-resources.html) to apply to the what if forecast.")
	forecast_createWhatIfForecastCmd.Flags().String("time-series-replacements-data-source", "", "The replacement time series dataset, which contains the rows that you want to change in the related time series dataset.")
	forecast_createWhatIfForecastCmd.Flags().String("time-series-transformations", "", "The transformations that are applied to the baseline time series.")
	forecast_createWhatIfForecastCmd.Flags().String("what-if-analysis-arn", "", "The Amazon Resource Name (ARN) of the what-if analysis.")
	forecast_createWhatIfForecastCmd.Flags().String("what-if-forecast-name", "", "The name of the what-if forecast.")
	forecast_createWhatIfForecastCmd.MarkFlagRequired("what-if-analysis-arn")
	forecast_createWhatIfForecastCmd.MarkFlagRequired("what-if-forecast-name")
	forecastCmd.AddCommand(forecast_createWhatIfForecastCmd)
}
