package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getCostForecastCmd = &cobra.Command{
	Use:   "get-cost-forecast",
	Short: "Retrieves a forecast for how much Amazon Web Services predicts that you will spend over the forecast time period that you select, based on your past costs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getCostForecastCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_getCostForecastCmd).Standalone()

		ce_getCostForecastCmd.Flags().String("billing-view-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies a specific billing view.")
		ce_getCostForecastCmd.Flags().String("filter", "", "The filters that you want to use to filter your forecast.")
		ce_getCostForecastCmd.Flags().String("granularity", "", "How granular you want the forecast to be.")
		ce_getCostForecastCmd.Flags().String("metric", "", "Which metric Cost Explorer uses to create your forecast.")
		ce_getCostForecastCmd.Flags().String("prediction-interval-level", "", "Cost Explorer always returns the mean forecast as a single point.")
		ce_getCostForecastCmd.Flags().String("time-period", "", "The period of time that you want the forecast to cover.")
		ce_getCostForecastCmd.MarkFlagRequired("granularity")
		ce_getCostForecastCmd.MarkFlagRequired("metric")
		ce_getCostForecastCmd.MarkFlagRequired("time-period")
	})
	ceCmd.AddCommand(ce_getCostForecastCmd)
}
