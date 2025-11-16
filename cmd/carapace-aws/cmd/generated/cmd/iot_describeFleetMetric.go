package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeFleetMetricCmd = &cobra.Command{
	Use:   "describe-fleet-metric",
	Short: "Gets information about the specified fleet metric.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeFleetMetricCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeFleetMetricCmd).Standalone()

		iot_describeFleetMetricCmd.Flags().String("metric-name", "", "The name of the fleet metric to describe.")
		iot_describeFleetMetricCmd.MarkFlagRequired("metric-name")
	})
	iotCmd.AddCommand(iot_describeFleetMetricCmd)
}
