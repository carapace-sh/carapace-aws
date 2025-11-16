package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getTimeSeriesDataPointCmd = &cobra.Command{
	Use:   "get-time-series-data-point",
	Short: "Gets the existing data point for the asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getTimeSeriesDataPointCmd).Standalone()

	datazone_getTimeSeriesDataPointCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain that houses the asset for which you want to get the data point.")
	datazone_getTimeSeriesDataPointCmd.Flags().String("entity-identifier", "", "The ID of the asset for which you want to get the data point.")
	datazone_getTimeSeriesDataPointCmd.Flags().String("entity-type", "", "The type of the asset for which you want to get the data point.")
	datazone_getTimeSeriesDataPointCmd.Flags().String("form-name", "", "The name of the time series form that houses the data point that you want to get.")
	datazone_getTimeSeriesDataPointCmd.Flags().String("identifier", "", "The ID of the data point that you want to get.")
	datazone_getTimeSeriesDataPointCmd.MarkFlagRequired("domain-identifier")
	datazone_getTimeSeriesDataPointCmd.MarkFlagRequired("entity-identifier")
	datazone_getTimeSeriesDataPointCmd.MarkFlagRequired("entity-type")
	datazone_getTimeSeriesDataPointCmd.MarkFlagRequired("form-name")
	datazone_getTimeSeriesDataPointCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_getTimeSeriesDataPointCmd)
}
