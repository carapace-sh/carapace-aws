package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateCustomMetricCmd = &cobra.Command{
	Use:   "update-custom-metric",
	Short: "Updates a Device Defender detect custom metric.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateCustomMetricCmd).Standalone()

	iot_updateCustomMetricCmd.Flags().String("display-name", "", "Field represents a friendly name in the console for the custom metric, it doesn't have to be unique.")
	iot_updateCustomMetricCmd.Flags().String("metric-name", "", "The name of the custom metric.")
	iot_updateCustomMetricCmd.MarkFlagRequired("display-name")
	iot_updateCustomMetricCmd.MarkFlagRequired("metric-name")
	iotCmd.AddCommand(iot_updateCustomMetricCmd)
}
