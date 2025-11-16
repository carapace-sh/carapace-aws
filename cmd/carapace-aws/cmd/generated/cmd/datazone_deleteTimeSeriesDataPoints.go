package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteTimeSeriesDataPointsCmd = &cobra.Command{
	Use:   "delete-time-series-data-points",
	Short: "Deletes the specified time series form for the specified asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteTimeSeriesDataPointsCmd).Standalone()

	datazone_deleteTimeSeriesDataPointsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
	datazone_deleteTimeSeriesDataPointsCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain that houses the asset for which you want to delete a time series form.")
	datazone_deleteTimeSeriesDataPointsCmd.Flags().String("entity-identifier", "", "The ID of the asset for which you want to delete a time series form.")
	datazone_deleteTimeSeriesDataPointsCmd.Flags().String("entity-type", "", "The type of the asset for which you want to delete a time series form.")
	datazone_deleteTimeSeriesDataPointsCmd.Flags().String("form-name", "", "The name of the time series form that you want to delete.")
	datazone_deleteTimeSeriesDataPointsCmd.MarkFlagRequired("domain-identifier")
	datazone_deleteTimeSeriesDataPointsCmd.MarkFlagRequired("entity-identifier")
	datazone_deleteTimeSeriesDataPointsCmd.MarkFlagRequired("entity-type")
	datazone_deleteTimeSeriesDataPointsCmd.MarkFlagRequired("form-name")
	datazoneCmd.AddCommand(datazone_deleteTimeSeriesDataPointsCmd)
}
