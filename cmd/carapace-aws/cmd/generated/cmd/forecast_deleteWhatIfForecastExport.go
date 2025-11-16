package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_deleteWhatIfForecastExportCmd = &cobra.Command{
	Use:   "delete-what-if-forecast-export",
	Short: "Deletes a what-if forecast export created using the [CreateWhatIfForecastExport]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_deleteWhatIfForecastExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_deleteWhatIfForecastExportCmd).Standalone()

		forecast_deleteWhatIfForecastExportCmd.Flags().String("what-if-forecast-export-arn", "", "The Amazon Resource Name (ARN) of the what-if forecast export that you want to delete.")
		forecast_deleteWhatIfForecastExportCmd.MarkFlagRequired("what-if-forecast-export-arn")
	})
	forecastCmd.AddCommand(forecast_deleteWhatIfForecastExportCmd)
}
