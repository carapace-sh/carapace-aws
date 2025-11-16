package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteFleetMetricCmd = &cobra.Command{
	Use:   "delete-fleet-metric",
	Short: "Deletes the specified fleet metric.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteFleetMetricCmd).Standalone()

	iot_deleteFleetMetricCmd.Flags().String("expected-version", "", "The expected version of the fleet metric to delete.")
	iot_deleteFleetMetricCmd.Flags().String("metric-name", "", "The name of the fleet metric to delete.")
	iot_deleteFleetMetricCmd.MarkFlagRequired("metric-name")
	iotCmd.AddCommand(iot_deleteFleetMetricCmd)
}
