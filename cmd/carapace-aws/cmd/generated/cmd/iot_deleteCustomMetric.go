package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteCustomMetricCmd = &cobra.Command{
	Use:   "delete-custom-metric",
	Short: "Deletes a Device Defender detect custom metric.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteCustomMetricCmd).Standalone()

	iot_deleteCustomMetricCmd.Flags().String("metric-name", "", "The name of the custom metric.")
	iot_deleteCustomMetricCmd.MarkFlagRequired("metric-name")
	iotCmd.AddCommand(iot_deleteCustomMetricCmd)
}
