package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getUsageForecastCmd = &cobra.Command{
	Use:   "get-usage-forecast",
	Short: "Retrieves a forecast for how much Amazon Web Services predicts that you will use over the forecast time period that you select, based on your past usage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getUsageForecastCmd).Standalone()

	ce_getUsageForecastCmd.Flags().String("billing-view-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies a specific billing view.")
	ce_getUsageForecastCmd.Flags().String("filter", "", "The filters that you want to use to filter your forecast.")
	ce_getUsageForecastCmd.Flags().String("granularity", "", "How granular you want the forecast to be.")
	ce_getUsageForecastCmd.Flags().String("metric", "", "Which metric Cost Explorer uses to create your forecast.")
	ce_getUsageForecastCmd.Flags().String("prediction-interval-level", "", "Amazon Web Services Cost Explorer always returns the mean forecast as a single point.")
	ce_getUsageForecastCmd.Flags().String("time-period", "", "The start and end dates of the period that you want to retrieve usage forecast for.")
	ce_getUsageForecastCmd.MarkFlagRequired("granularity")
	ce_getUsageForecastCmd.MarkFlagRequired("metric")
	ce_getUsageForecastCmd.MarkFlagRequired("time-period")
	ceCmd.AddCommand(ce_getUsageForecastCmd)
}
