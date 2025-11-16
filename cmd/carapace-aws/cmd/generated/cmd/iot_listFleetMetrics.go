package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listFleetMetricsCmd = &cobra.Command{
	Use:   "list-fleet-metrics",
	Short: "Lists all your fleet metrics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listFleetMetricsCmd).Standalone()

	iot_listFleetMetricsCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
	iot_listFleetMetricsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise `null` to receive the first set of results.")
	iotCmd.AddCommand(iot_listFleetMetricsCmd)
}
