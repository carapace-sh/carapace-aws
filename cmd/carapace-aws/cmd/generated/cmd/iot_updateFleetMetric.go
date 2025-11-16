package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateFleetMetricCmd = &cobra.Command{
	Use:   "update-fleet-metric",
	Short: "Updates the data for a fleet metric.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateFleetMetricCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateFleetMetricCmd).Standalone()

		iot_updateFleetMetricCmd.Flags().String("aggregation-field", "", "The field to aggregate.")
		iot_updateFleetMetricCmd.Flags().String("aggregation-type", "", "The type of the aggregation query.")
		iot_updateFleetMetricCmd.Flags().String("description", "", "The description of the fleet metric.")
		iot_updateFleetMetricCmd.Flags().String("expected-version", "", "The expected version of the fleet metric record in the registry.")
		iot_updateFleetMetricCmd.Flags().String("index-name", "", "The name of the index to search.")
		iot_updateFleetMetricCmd.Flags().String("metric-name", "", "The name of the fleet metric to update.")
		iot_updateFleetMetricCmd.Flags().String("period", "", "The time in seconds between fleet metric emissions.")
		iot_updateFleetMetricCmd.Flags().String("query-string", "", "The search query string.")
		iot_updateFleetMetricCmd.Flags().String("query-version", "", "The version of the query.")
		iot_updateFleetMetricCmd.Flags().String("unit", "", "Used to support unit transformation such as milliseconds to seconds.")
		iot_updateFleetMetricCmd.MarkFlagRequired("index-name")
		iot_updateFleetMetricCmd.MarkFlagRequired("metric-name")
	})
	iotCmd.AddCommand(iot_updateFleetMetricCmd)
}
