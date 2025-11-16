package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_createMonitorCmd = &cobra.Command{
	Use:   "create-monitor",
	Short: "Creates a predictor monitor resource for an existing auto predictor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_createMonitorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_createMonitorCmd).Standalone()

		forecast_createMonitorCmd.Flags().String("monitor-name", "", "The name of the monitor resource.")
		forecast_createMonitorCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the predictor to monitor.")
		forecast_createMonitorCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/forecast/latest/dg/tagging-forecast-resources.html) to apply to the monitor resource.")
		forecast_createMonitorCmd.MarkFlagRequired("monitor-name")
		forecast_createMonitorCmd.MarkFlagRequired("resource-arn")
	})
	forecastCmd.AddCommand(forecast_createMonitorCmd)
}
