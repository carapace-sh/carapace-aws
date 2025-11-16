package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createCustomMetricCmd = &cobra.Command{
	Use:   "create-custom-metric",
	Short: "Use this API to define a Custom Metric published by your devices to Device Defender.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createCustomMetricCmd).Standalone()

	iot_createCustomMetricCmd.Flags().String("client-request-token", "", "Each custom metric must have a unique client request token.")
	iot_createCustomMetricCmd.Flags().String("display-name", "", "The friendly name in the console for the custom metric.")
	iot_createCustomMetricCmd.Flags().String("metric-name", "", "The name of the custom metric.")
	iot_createCustomMetricCmd.Flags().String("metric-type", "", "The type of the custom metric.")
	iot_createCustomMetricCmd.Flags().String("tags", "", "Metadata that can be used to manage the custom metric.")
	iot_createCustomMetricCmd.MarkFlagRequired("client-request-token")
	iot_createCustomMetricCmd.MarkFlagRequired("metric-name")
	iot_createCustomMetricCmd.MarkFlagRequired("metric-type")
	iotCmd.AddCommand(iot_createCustomMetricCmd)
}
