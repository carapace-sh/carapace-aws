package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createFleetMetricCmd = &cobra.Command{
	Use:   "create-fleet-metric",
	Short: "Creates a fleet metric.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createFleetMetricCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createFleetMetricCmd).Standalone()

		iot_createFleetMetricCmd.Flags().String("aggregation-field", "", "The field to aggregate.")
		iot_createFleetMetricCmd.Flags().String("aggregation-type", "", "The type of the aggregation query.")
		iot_createFleetMetricCmd.Flags().String("description", "", "The fleet metric description.")
		iot_createFleetMetricCmd.Flags().String("index-name", "", "The name of the index to search.")
		iot_createFleetMetricCmd.Flags().String("metric-name", "", "The name of the fleet metric to create.")
		iot_createFleetMetricCmd.Flags().String("period", "", "The time in seconds between fleet metric emissions.")
		iot_createFleetMetricCmd.Flags().String("query-string", "", "The search query string.")
		iot_createFleetMetricCmd.Flags().String("query-version", "", "The query version.")
		iot_createFleetMetricCmd.Flags().String("tags", "", "Metadata, which can be used to manage the fleet metric.")
		iot_createFleetMetricCmd.Flags().String("unit", "", "Used to support unit transformation such as milliseconds to seconds.")
		iot_createFleetMetricCmd.MarkFlagRequired("aggregation-field")
		iot_createFleetMetricCmd.MarkFlagRequired("aggregation-type")
		iot_createFleetMetricCmd.MarkFlagRequired("metric-name")
		iot_createFleetMetricCmd.MarkFlagRequired("period")
		iot_createFleetMetricCmd.MarkFlagRequired("query-string")
	})
	iotCmd.AddCommand(iot_createFleetMetricCmd)
}
