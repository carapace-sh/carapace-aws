package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listMetricValuesCmd = &cobra.Command{
	Use:   "list-metric-values",
	Short: "Lists the values reported for an IoT Device Defender metric (device-side metric, cloud-side metric, or custom metric) by the given thing during the specified time period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listMetricValuesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listMetricValuesCmd).Standalone()

		iot_listMetricValuesCmd.Flags().String("dimension-name", "", "The dimension name.")
		iot_listMetricValuesCmd.Flags().String("dimension-value-operator", "", "The dimension value operator.")
		iot_listMetricValuesCmd.Flags().String("end-time", "", "The end of the time period for which metric values are returned.")
		iot_listMetricValuesCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iot_listMetricValuesCmd.Flags().String("metric-name", "", "The name of the security profile metric for which values are returned.")
		iot_listMetricValuesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		iot_listMetricValuesCmd.Flags().String("start-time", "", "The start of the time period for which metric values are returned.")
		iot_listMetricValuesCmd.Flags().String("thing-name", "", "The name of the thing for which security profile metric values are returned.")
		iot_listMetricValuesCmd.MarkFlagRequired("end-time")
		iot_listMetricValuesCmd.MarkFlagRequired("metric-name")
		iot_listMetricValuesCmd.MarkFlagRequired("start-time")
		iot_listMetricValuesCmd.MarkFlagRequired("thing-name")
	})
	iotCmd.AddCommand(iot_listMetricValuesCmd)
}
